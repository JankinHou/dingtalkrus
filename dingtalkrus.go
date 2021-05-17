package dingtalkrus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type DingTalkHook struct {
	c http.Client
}

type requestMessageText struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func NewDingTalkHook(opts ...Option) (*DingTalkHook, error) {
	mergeOptions(opts...)
	checkOptions()
	return &DingTalkHook{
		c: http.Client{},
	}, nil
}

// Fire is called when a log event is fired.
func (hook *DingTalkHook) Fire(entry *logrus.Entry) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	request := requestMessageText{}

	request.At.AtMobiles = options.AtMobiles
	request.At.AtUserIds = options.AtUserIds
	request.At.IsAtAll = options.IsAtAll
	request.Msgtype = string(options.MsgType)
	request.Text.Content = getMessage(entry)

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	// https://oapi.dingtalk.com/robot/send?access_token=1f72fe3c0c08490464da3096253ce2f3942184ebb1dcf259d5f69a33bc3de12d
	urlStr := "https://oapi.dingtalk.com/robot/send?access_token=" + options.Access_token
	req, err := HttpRequest(ctx, urlStr, "POST", string(requestJSON))

	if err != nil {
		return err

	}

	var response struct {
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	json.Unmarshal(req, &response)
	if response.Errcode == 0 && response.Errmsg == "ok" {
		return nil
	}

	return fmt.Errorf("error code %d with message: %s", response.Errcode, response.Errmsg)
}

// Levels returns the available logging levels.
func (hook *DingTalkHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}

func getMessage(entry *logrus.Entry) string {
	message := options.MessageFormat.ToString()
	message = strings.ReplaceAll(message, "{{app}}", options.AppName)
	message = strings.ReplaceAll(message, "{{time}}", entry.Time.In(options.TimeZone).Format(options.TimeFormat))
	message = strings.ReplaceAll(message, "{{level}}", entry.Level.String())
	message = strings.ReplaceAll(message, "{{message}}", entry.Message)

	fields, _ := json.MarshalIndent(entry.Data, "", "\t")
	message = strings.ReplaceAll(message, "{{content}}", string(fields))
	return message
}
