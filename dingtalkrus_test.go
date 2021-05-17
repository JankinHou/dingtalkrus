package dingtalkrus

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestDingTalkRus(t *testing.T) {

	hook, err := NewDingTalkHook(
		Option{
			AppName:       "探日工程",
			MessageFormat: nil,
			TimeFormat:    "",
			TimeZone:      &time.Location{},
			MsgType:       "text",
			AtMobiles:     []string{},
			AtUserIds:     []string{},
			IsAtAll:       false,
			Access_token:  "",
			Keywords:      "logrus",
		},
	)
	if err != nil {
		t.Error(err.Error())
	}
	logr := logrus.New()
	logr.Hooks.Add(hook)

	logr.WithFields(logrus.Fields{
		"t1": "a1",
	}).Warn("asass")
	// logr.Error("123456")
	t.Error("done")

}
