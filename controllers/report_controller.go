package contollers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GenerateAndPrint(incoming, data string) {
	var numberOfPrint = 1
	tmp := strings.Split(incoming, " ")
	if len(tmp) == 3 {
		number, err := strconv.Atoi(tmp[2])
		if err != nil {
			log.Println("convert string to int err: ", err)
		}
		numberOfPrint = number
	}
	log.Println(numberOfPrint)
	//example type
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	//create a new xlsx file and write a struct
	//in a new row
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet.SetColWidth(3, 3, 5)
	style := xlsx.NewStyle()
	style.Alignment = xlsx.Alignment{
		Horizontal: "left",
		Vertical:   "center",
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	for i := 0; i < 7; i++ {
		cell = sheet.Cell(i*6, 0)
		cell.SetStyle(style)
		cell.Merge(2, 3)
		cell.Value = data
		cell = sheet.Cell(i*6, 4)
		cell.SetStyle(style)
		cell.Merge(2, 3)
		cell.Value = data
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
