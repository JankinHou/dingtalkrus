# dingtalkrus
## Intro


## Installation
Install the package with go:
```go
go get github.com/JankinHou/dingtalkrus
```

## Usage
```go
import (
    "github.com/sirupsen/logrus"
    "github.com/JankinHou/dingtalkrus"
)

func main() {
  log       := logrus.New()
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
	if err == nil {
		log.Hooks.Add(hook)
	}
}
```


## License
MIT