package tgbotclient

import (
	"encoding/json"
	"errors"
	"github.com/annamuratov1987/tgbotclient/pkg/domain"
	"io"
	"net/http"
	"net/url"
)

var (
	BaseUrl = url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
	}
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(token string) *Client {
	return &Client{
		client: http.DefaultClient,
		token:  token,
	}
}

func (c *Client) getMethodUrl(method string) *url.URL {
	return BaseUrl.JoinPath("bot"+c.token, method)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	return resp, err
}

func LoadResult[T any](response *http.Response, result *T) error {
	var encResp Response[T]

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &encResp)

	if encResp.Ok {
		*result = encResp.Result
	} else if encResp.Description != "" {
		return errors.New(encResp.Description)
	}

	return err
}

func (c *Client) GetUpdates(builder RequestBuilder) ([]domain.Update, error) {
	url := c.getMethodUrl("getUpdates")

	req, err := builder.BuildRequest(*url)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var result []domain.Update
	err = LoadResult(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) SendMessage(builder RequestBuilder) (result domain.Message, err error) {
	url := c.getMethodUrl("sendMessage")

	req, err := builder.BuildRequest(*url)
	if err != nil {
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	err = LoadResult(resp, &result)

	return
}

func (c *Client) SendPhoto(builder RequestBuilder) (result domain.Message, err error) {
	url := c.getMethodUrl("sendPhoto")

	req, err := builder.BuildRequest(*url)
	if err != nil {
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = LoadResult(resp, &result)

	return
}
