package models

import "gorm.io/gorm"

var KeywordSplitCreate = []string{"เพิ่ม", "อำเภอ", "จังหวัด", "ผู้ส่ง", "~"}

var KeywordSplitShow = []string{"แสดง", "~"}

var KeywordSplitPrint = []string{"ปริ้น", "จำนวน", "~"}

var KeywordSplitDelete = []string{"ลบ", "~"}

type Customers struct {
	gorm.Model
	Name     string `gorm:"name"`
	District string `gorm:"district"`
	Province string `gorm:"province"`
	Sender   string `gorm:"sender"`
}
