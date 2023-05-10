package tgbotclient

import (
	"net/http"
	"net/url"
)

type RequestBuilder interface {
	BuildRequest(url url.URL) (*http.Request, error)
}
