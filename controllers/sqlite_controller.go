package contollers

import (
	"fmt"
	"log"
	"shipSlip/models"
	"shipSlip/services"
	"strings"
)

var sqliteClient services.Sqlite

func ConnectToSqlite() {
	var err error
	sqliteClient, err = services.Connect("shipSlip.db")
	if err != nil {
		log.Panic("Connect to db err: ", err)
	}
}

func CreateCustomer(data string) bool {
	log.Println("Insert customer: ", data)
	tmp := strings.SplitAfter(data, "(")
	tmp = strings.Split(tmp[1], ")")
	tmp = strings.Split(tmp[0], ".")
	if err := sqliteClient.Insert(&models.Customers{
		Name:     tmp[0],
		District: tmp[1],
		Province: tmp[2],
		Sender:   tmp[3],
	}); err != nil {
		log.Println("Insert customer to db err: ", err)
		return false
	}
	return true
}

func GetAllCustomers() string {
	var customers []models.Customers
	var customerStr string
	if err := sqliteClient.QueryAll(&customers); err != nil {
		log.Println("Query all customer err: ", err)
	}
	customerStr = fmt.Sprintf("รายชื่อร้านค้าทั้งหมด \n")
	for index, value := range customers {
		customerStr += fmt.Sprintf("%d. ร้าน%s อำเภอ%s จังหวัด%s ผู้ส่ง%s\n", index+1, value.Name, value.District, value.Province, value.Sender)
	}
	return customerStr
}
