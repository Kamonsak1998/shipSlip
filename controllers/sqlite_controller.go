package contollers

import (
	"log"
	"shipSlip/models"
	"shipSlip/services"
)

var sqliteClient services.Sqlite

func ConnectToSqlite() {
	var err error
	sqliteClient, err = services.Connect("shipSlip.db")
	if err != nil {
		log.Panic("Connect to db err: ", err)
	}
}

func InsertCustomer(data string) {
	log.Println("Insert customer: ", data)
	if err := sqliteClient.Insert(&models.Customers{Name: "ทดสอบ", District: "บางขุนเทียน", Province: "กรุงเทพ", Sender: "ทินกร จงวิไลเกษม"}); err != nil {
		log.Println("Insert customer to db err: ", err)
	}
}
