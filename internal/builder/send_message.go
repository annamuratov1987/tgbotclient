package builder

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SendMessageRequestBuilder struct {
	BuilderOption
}

func NewSendMessageRequestBuilder() *SendMessageRequestBuilder {
	return &SendMessageRequestBuilder{map[string]interface{}{}}
}

func (b SendMessageRequestBuilder) ValidateOption() (bool, error) {
	if val := b.GetOption("chat_id"); val != nil {
		_, intOk := val.(int64)
		_, strOk := val.(string)
		if !intOk && !strOk {
			return false, fmt.Errorf("'%s': %w", "chat_id", ErrorMismatchTypeStringOrInt64)
		}
	} else {
		return false, fmt.Errorf("'%s': %w", "chat_id", ErrorRequired)
	}

	if val := b.GetOption("text"); val != nil {
		if _, ok := val.(string); !ok {
			return false, fmt.Errorf("'%s': %w", "text", ErrorMismatchTypeString)
		}
	} else {
		return false, fmt.Errorf("'%s': %w", "text", ErrorRequired)
	}

	return true, nil
}

func (b SendMessageRequestBuilder) BuildRequest(url url.URL) (*http.Request, error) {

	if ok, err := b.ValidateOption(); !ok {
		return nil, err
	}

	q := url.Query()

	val := b.GetOption("chat_id")
	if num, ok := val.(int64); ok {
		q.Add("chat_id", strconv.FormatInt(num, 10))
	}
	if str, ok := val.(string); ok {
		q.Add("chat_id", str)
	}

	q.Add("text", b.GetOption("text").(string))

	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
