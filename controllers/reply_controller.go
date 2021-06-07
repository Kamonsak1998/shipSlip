package contollers

import (
	"log"
	linbotService "shipSlip/services"
)

func ReplyMessage(replyToken, replyMsg string) {
	if err := linbotService.ReplyMessage(replyToken, replyMsg); err != nil {
		log.Print("reply message err:", err)
	}
}
