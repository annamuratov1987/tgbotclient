package domain

type Message struct {
	Id   int64  `json:"message_id"`
	Date int64  `json:"date"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text,omitempty"`
}
