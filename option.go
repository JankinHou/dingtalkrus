package dingtalkrus

import (
	"time"
)

type SafeSwitcher bool

const (
	// SafeOn enables the message safe
	SafeOn SafeSwitcher = true
	// SafeOff disenables the message safe
	SafeOff SafeSwitcher = false
)

// ToInt for translating to int
func (s SafeSwitcher) ToInt() int {
	if s {
		return 1
	}

	return 0
}

// MessageType defines ...
type MessageType string

const (
	// TextMessage for sending text message
	TextMessage MessageType = "text"
)

type MessageFormater interface {
	ToString() string
}

type normalTextFormat struct {
	text string
}

func (n normalTextFormat) ToString() string {
	return n.text
}

type Option struct {
	// AppName will be used for logging out
	AppName string
	// MessageFormart defines the format of the message
	// It is an interface that you can customize the output as long as you implement it
	MessageFormat MessageFormater
	// TimeFormat defines ...
	TimeFormat string
	// TimeZone defines ...
	TimeZone *time.Location
	// Safe defines ...
	Safe SafeSwitcher
	// MsgType defines what type of log you want to display in WeCom
	MsgType      MessageType
	AtMobiles    []string
	AtUserIds    []string
	IsAtAll      bool
	Access_token string
	Keywords     string
}

var options Option

func checkOptions() {

	if options.AppName == "" {
		options.AppName = "Undefined"
	}
	if options.TimeZone == nil {
		tz, err := time.LoadLocation("Asia/Chongqing")
		if err != nil {
			panic(err)
		}
		options.TimeZone = tz
	}
	if options.TimeFormat == "" {
		options.TimeFormat = "01-02 15:04:05"
	}
	if options.MessageFormat == nil {
		var normalMessage normalTextFormat

		normalMessage.text = options.Keywords + ` Monitor
***********
* AppName: {{app}}
* Time: {{time}}
* Level: {{level}}
***********
* Message: {{message}}
***********
{{content}}`
		options.MessageFormat = normalMessage
	}

}

func mergeOptions(opts ...Option) {
	// Set defaults
	options.MsgType = TextMessage

	for _, opt := range opts {
		if opt.AppName != "" {
			options.AppName = opt.AppName
		}
		if opt.MessageFormat != nil {
			options.MessageFormat = opt.MessageFormat
		}

		if opt.TimeFormat != "" {
			options.TimeFormat = opt.TimeFormat
		}
		if opt.TimeZone != nil {
			options.TimeZone = opt.TimeZone
		}

		if opt.MsgType != TextMessage {
			options.MsgType = opt.MsgType
		}
		if opt.Access_token != "" {
			options.Access_token = opt.Access_token
		}
		if opt.Keywords != "" {
			options.Keywords = opt.Keywords
		}
		if len(opt.AtMobiles) > 0 {
			options.AtMobiles = opt.AtMobiles
		}
		if len(opt.AtUserIds) > 0 {
			options.AtUserIds = opt.AtUserIds
		}
		options.IsAtAll = opt.IsAtAll

	}
}
