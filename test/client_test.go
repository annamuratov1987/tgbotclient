package test

import (
	"fmt"
	"os"
	"testing"
	"tgbotclient"
	"tgbotclient/internal/builder"
)

var token = "5750139837:AAHQzej9sL-bL5b7EzZWdaeJ9fv1ak-PtFw"

func Test_ClientGetUpdates(t *testing.T) {
	client := tgbotclient.NewClient(token)

	var builder = builder.NewGetUpdateRequestBuilder()

	updates, err := client.GetUpdates(builder)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(updates)
}

func Test_ClientSendMessage(t *testing.T) {
	client := tgbotclient.NewClient(token)

	builder := builder.NewSendMessageRequestBuilder()
	builder.SetOption("chat_id", int64(1283280061))
	builder.SetOption("text", "Hello")

	_, err := client.SendMessage(builder)
	if err != nil {
		t.Error(err)
	}
}

func Test_ClientSendPhoto(t *testing.T) {
	client := tgbotclient.NewClient(token)

	builder := builder.NewSendPhotoRequestBuilder()
	builder.SetOption("chat_id", int64(1283280061))

	photo, err := os.Open("avto.png")
	if err != nil {
		t.Error(err)
	}

	builder.SetOption("photo", *photo)

	_, err = client.SendPhoto(builder)
	if err != nil {
		t.Error(err)
	}
}
