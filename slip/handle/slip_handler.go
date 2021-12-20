package handler

import (
	"fmt"
	"log"
	contollers "shipSlip/controllers"
	linbotControllers "shipSlip/controllers"
	"shipSlip/models"
	linbotService "shipSlip/services"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	chSecret = "9bcd90c7cf33f1e8c6e9f1052fbb3476"
	chAccess = "gTdR5OhomqGGsFoJCewZz0Oo3xL2OseGFgXT/x0vvZVEA9bHbFzwCLW6sT1sN8jm1b2tAiOGpOFmCgz48DGrlaxHOVUxCFZrV5cQyv7qWctC0mf+MpmGSHYLvSy7bHhT1b8/2SuLAMhmmzf6cCnDCQdB04t89/1O/w1cDnyilFU="
)

// Handler from Messaging API
func Handler(ctx echo.Context) error {
	linbotService.New(chSecret, chAccess)
	events, err := linbotService.ParseRequestToEvents(ctx.Request())
	if err != nil {
		return err
	}
	for _, event := range events {
		if event.Type == "message" {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if strings.HasPrefix(message.Text, models.KeywordSplitCreate[0]) {
					log.Println("เพิ่มร้าน->", message.Text)
					ok := contollers.CreateCustomer(message.Text)
					if ok {
						linbotControllers.ReplyMessage(event.ReplyToken, "เพิ่มร้านค้า \"สำเร็จ\"")
					} else {
						linbotControllers.ReplyMessage(event.ReplyToken, "เพิ่มร้านค้า \"ไม่สำเร็จ\"")
					}
				} else if strings.HasPrefix(message.Text, "แสดงทั้งหมด") {
					log.Println("แสดงร้านทั้งหมด->", message.Text)
					customers := contollers.GetAllCustomers()
					linbotControllers.ReplyMessage(event.ReplyToken, customers)
				} else if strings.HasPrefix(message.Text, models.KeywordSplitShow[0]) {
					log.Println("แสดงร้าน->", message.Text)
					customer, ok := linbotControllers.GetCustomer(models.KeywordSplitShow, message.Text)
					if ok {
						linbotControllers.ReplyMessage(event.ReplyToken, customer)
					} else {
						linbotControllers.ReplyMessage(event.ReplyToken, "ไม่มีข้อมูลของร้าน \""+customer+"\"")
					}
				} else if strings.HasPrefix(message.Text, models.KeywordSplitDelete[0]) {
					log.Println("ลบร้าน->", message.Text)
					ok := linbotControllers.DeleteCustomer(message.Text)
					if ok {
						linbotControllers.ReplyMessage(event.ReplyToken, "ลบร้านค้า \"สำเร็จ\"")
					} else {
						linbotControllers.ReplyMessage(event.ReplyToken, "ลบร้านค้า \"ไม่สำเสร็จ\"")
					}
				} else if strings.HasPrefix(message.Text, models.KeywordSplitPrint[0]) {
					log.Println("ปริ้น->", message.Text)
					customer, ok := linbotControllers.GetCustomer(models.KeywordSplitPrint, message.Text)
					if ok {
						linbotControllers.GenerateAndPrint(message.Text, customer)
						linbotControllers.ReplyMessage(event.ReplyToken, "ปริ้นร้านค้า \"สำเร็จ\"")
					} else {
						linbotControllers.ReplyMessage(event.ReplyToken, "ปริ้นร้านค้า \""+customer+"\" ไม่สำเร็จ")
					}
				} else {
					log.Print("default->", message.Text)
					replyMsg := "กรุณาเลือกคำสั่งตามด้านล่าง:\n-เพิ่ม...\n-แสดงทั้งหมด\n-แสดง...\n-ลบ...\n-ปริ้น..."
					linbotControllers.ReplyMessage(event.ReplyToken, replyMsg)
				}

			case *linebot.StickerMessage:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
				if err := linbotService.ReplyMessage(event.ReplyToken, replyMessage); err != nil {
					log.Print("reply message err:", err)
				}
			}
		}
	}
	return ctx.JSON(200, "")
}
