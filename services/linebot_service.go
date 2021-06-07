package services

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	linebotClient  *linebot.Client
	TextMessage    linebot.TextMessage
	StickerMessage linebot.StickerMessage
)

func New(chSecret, chAccess string) error {
	var err error
	client := &http.Client{}
	linebotClient, err = linebot.New(chSecret, chAccess, linebot.WithHTTPClient(client))
	if err != nil {
		return err
	}
	return nil
}

func ParseRequestToEvents(req *http.Request) ([]*linebot.Event, error) {
	return linebotClient.ParseRequest(req)
}

func ReplyMessage(replyToken, replyMessage string) error {
	if _, err := linebotClient.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}
