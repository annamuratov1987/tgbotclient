package builder

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type SendPhotoRequestBuilder struct {
	BuilderOption
}

func NewSendPhotoRequestBuilder() *SendPhotoRequestBuilder {
	return &SendPhotoRequestBuilder{map[string]interface{}{}}
}

func (b SendPhotoRequestBuilder) ValidateOption() (bool, error) {
	if val := b.GetOption("chat_id"); val != nil {
		_, intOk := val.(int64)
		_, strOk := val.(string)
		if !intOk && !strOk {
			return false, fmt.Errorf("'%s': %w", "chat_id", ErrorMismatchTypeStringOrInt64)
		}
	} else {
		return false, fmt.Errorf("'%s': %w", "chat_id", ErrorRequired)
	}

	if val := b.GetOption("photo"); val != nil {
		_, fileOk := val.(os.File)
		_, strOk := val.(string)
		if !fileOk && !strOk {
			return false, fmt.Errorf("'%s': %w", "photo", ErrorMismatchTypeFileOrInt64)
		}
	} else {
		return false, fmt.Errorf("'%s': %w", "photo", ErrorRequired)
	}

	return true, nil
}

func (b SendPhotoRequestBuilder) BuildRequest(url url.URL) (*http.Request, error) {

	if ok, err := b.ValidateOption(); !ok {
		return &http.Request{}, err
	}

	q := url.Query()

	chatId := b.GetOption("chat_id")
	if num, ok := chatId.(int64); ok {
		q.Add("chat_id", strconv.FormatInt(num, 10))
	}
	if str, ok := chatId.(string); ok {
		q.Add("chat_id", str)
	}

	photo := b.GetOption("photo")
	if str, ok := photo.(string); ok {
		q.Add("photo", str)
	}

	url.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", url.String(), nil)

	if inputFile, ok := photo.(os.File); ok {
		r, w := io.Pipe()
		m := multipart.NewWriter(w)

		go func() {
			defer w.Close()
			defer m.Close()

			part, err := m.CreateFormFile("photo", inputFile.Name())
			if err != nil {
				w.CloseWithError(err)
				return
			}

			if _, err := io.Copy(part, &inputFile); err != nil {
				w.CloseWithError(err)
				return
			}
		}()

		req.Body = r
		req.Header.Set("Content-Type", m.FormDataContentType())
	}

	return req, err
}
