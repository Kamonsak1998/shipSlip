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
	customerData := ExtractToCustomer(models.KeywordSplitCreate, data)
	// if len(tmp) < 4 {
	// 	log.Printf("Insert customer to db err: %v, data: %s", fmt.Errorf("data not complete"), data)
	// 	return false
	// }
	if err := sqliteClient.Insert(customerData); err != nil {
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

func GetCustomer(keyword []string, name string) (string, bool) {
	customerData := ExtractToCustomer(keyword, name)

	var customers models.Customers
	var customerStr string
	if err := sqliteClient.Query(&customers, &models.Customers{Name: customerData.Name}); err != nil {
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
	customerData := ExtractToCustomer(models.KeywordSplitDelete, name)

	var customers models.Customers
	RowsAffected, err := sqliteClient.Delete(&customers, &models.Customers{Name: customerData.Name})
	if RowsAffected == 0 || err != nil {
		log.Println("Delete customer err: ", err)
		return false
	}
	return true
}

func ExtractToCustomer(keyword []string, data string) *models.Customers {
	customerData := &models.Customers{}
	if tmp := strings.Split(data, keyword[0])[1]; tmp != "" {
		tmp2 := strings.Split(tmp, keyword[1])
		customerData.Name = tmp2[0]
		if len(tmp2) == 1 {
			return customerData
		}
		log.Println("customerData.Name ", customerData.Name)
		tmp2 = strings.Split(tmp2[1], keyword[2])
		customerData.District = tmp2[0]
		if len(tmp2) == 1 {
			return customerData
		}
		tmp2 = strings.Split(tmp2[1], keyword[3])
		customerData.Sender = tmp2[1]
	}
	return customerData
}
