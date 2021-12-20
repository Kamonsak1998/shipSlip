package contollers

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GenerateAndPrint(incoming, data string) {
	numberOfPrint := 1
	tmp := strings.Split(incoming, "จำนวน")
	if len(tmp) == 2 {
		number, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Println("convert string to int err: ", err)
		}
		numberOfPrint = number
	}
	log.Println(numberOfPrint)
	// example type
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	// create a new xlsx file and write a struct
	// in a new row
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("ใบขนส่ง")
	if err != nil {
		fmt.Println(err.Error())
	}
	sheet.SetColWidth(3, 3, 5)
	style := xlsx.NewStyle()
	style.Alignment = xlsx.Alignment{
		Horizontal: "left",
		Vertical:   "center",
	}
	style.ApplyFont = true
	row = sheet.AddRow()
	cell = row.AddCell()
	for i := 0; i < 6*numberOfPrint; i++ {
		cell = sheet.Cell(i*6, 0)
		cell.Row.SetHeight(50)
		cell.SetStyle(style)
		cell.Merge(2, 3)
		cell.Value = data
		cell = sheet.Cell(i*6, 4)
		cell.Row.SetHeight(50)
		cell.SetStyle(style)
		cell.Merge(2, 3)
		cell.Value = data
	}

	fileName := "tmp"
	err = file.Save(fileName + ".xls")
	if err != nil {
		fmt.Println(err.Error())
	}

	filePDF := convertToPDF(fileName)
	print(filePDF)
}

func convertToPDF(fileName string) string {
	cmd := exec.Command("/bin/bash", "-c", "libreoffice --headless --invisible --convert-to pdf "+fileName+".xls")
	_, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	return fileName + ".pdf"
}

func print(fileName string) {
	cmd := exec.Command("/bin/bash", "-c", "lpr -U god -P LQ-310 "+fileName)
	_, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
}
