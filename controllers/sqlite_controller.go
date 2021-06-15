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
	if len(tmp) < 4 {
		log.Printf("Insert customer to db err: %v, data: %s", fmt.Errorf("data not complete"), data)
		return false
	}
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

func GetCustomer(name string) (string, bool) {
	name = strings.Split(name, " ")[1]
	var customers models.Customers
	var customerStr string
	if err := sqliteClient.Query(&customers, &models.Customers{Name: name}); err != nil {
		log.Println("Query customer err: ", err)
	}
	if customers.Name == "" {
		customerStr = fmt.Sprintf("ข้อมูลร้าน %s\n", customers.Name)
		return name, false
	}
	customerStr = fmt.Sprintf("ร้าน %s\n", customers.Name)
	customerStr += fmt.Sprintf("อำเภอ %s\n", customers.District)
	customerStr += fmt.Sprintf("จังหวัด %s\n", customers.Province)
	customerStr += fmt.Sprintf("ผู้ส่ง %s\n", customers.Sender)
	return customerStr, true
}

func DeleteCustomer(name string) bool {
	name = strings.Split(name, " ")[1]
	var customers models.Customers
	RowsAffected, err := sqliteClient.Delete(&customers, &models.Customers{Name: name})
	if RowsAffected == 0 || err != nil {
		log.Println("Delete customer err: ", err)
		return false
	}
	return true
}
