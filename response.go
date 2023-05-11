package tgbotclient

import "github.com/annamuratov1987/tgbotclient/pkg/domain"

type Response[ResultType any] struct {
	Ok          bool       `json:"ok"`
	Result      ResultType `json:"result,omitempty"`
	Description string     `json:"description,omitempty"`
	ErrorCode   int64      `json:"error_code,omitempty"`
}

type GetUpdatesResponse struct {
	Response[[]domain.Update]
}

type SendMessageResponse struct {
	Response[domain.Message]
}

type SendPhotoResponse struct {
	Response[domain.Message]
}
