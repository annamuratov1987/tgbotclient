package builder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type GetUpdateRequestBuilder struct {
	BuilderOption
}

func NewGetUpdateRequestBuilder() *GetUpdateRequestBuilder {
	return &GetUpdateRequestBuilder{map[string]interface{}{}}
}

func (b GetUpdateRequestBuilder) ValidateOption() (bool, error) {
	if val := b.GetOption("offset"); val != nil {
		if _, ok := val.(int64); !ok {
			return false, fmt.Errorf("'%s': %w", "offset", ErrorMismatchTypeInt64)
		}
	}

	if val := b.GetOption("limit"); val != nil {
		if _, ok := val.(int64); !ok {
			return false, fmt.Errorf("'%s': %w", "limit", ErrorMismatchTypeInt64)
		}
	}

	if val := b.GetOption("timeout"); val != nil {
		if _, ok := val.(int64); !ok {
			return false, fmt.Errorf("'%s': %w", "timeout", ErrorMismatchTypeInt64)
		}
	}

	if val := b.GetOption("allowed_updates"); val != nil {
		if _, ok := val.([]string); !ok {
			return false, fmt.Errorf("'%s': %w", "allowed_updates", ErrorMismatchTypeStringArray)
		}
	}

	return true, nil
}

func (b GetUpdateRequestBuilder) BuildRequest(url url.URL) (*http.Request, error) {
	if ok, err := b.ValidateOption(); !ok {
		return nil, err
	}

	q := url.Query()

	if val := b.GetOption("offset"); val != nil {
		q.Add("offset", strconv.FormatInt(val.(int64), 10))
	}

	if val := b.GetOption("limit"); val != nil {
		q.Add("limit", strconv.FormatInt(val.(int64), 10))
	}

	if val := b.GetOption("timeout"); val != nil {
		q.Add("timeout", strconv.FormatInt(val.(int64), 10))
	}

	if val := b.GetOption("allowed_updates"); val != nil {
		allowedUpdates, _ := json.Marshal(val.([]string))
		q.Add("allowed_updates", string(allowedUpdates))
	}

	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
