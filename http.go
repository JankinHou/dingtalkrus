package dingtalkrus

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func HttpRequest(ctx context.Context, urlStr, method, params string) (body []byte, err error) {
	fmt.Println(urlStr)
	req, err := http.NewRequestWithContext(ctx, method, urlStr, strings.NewReader(params))

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {

		return
	}
	return
}
