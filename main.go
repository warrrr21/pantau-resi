package main

import (
	"fmt"
)

func main() {

	listResi := ReadListResi()

	// loop list resi
	for i := range listResi.Resi {
		if listResi.Resi[i].Status != "DELIVERED" {
			CheckResi(listResi.Resi[i].Link)
		}
	}

}

func CheckResi(url string) {

	doc := GetHtml(url)
	data := ParseHtml(doc)

	filename := fmt.Sprintf("history/%s.json", data.NoResi)

	isFileExist := DoesFileExist(filename)

	if !isFileExist {
		fmt.Println("File Not Exist")
		SendUpdateResi(data)
		SaveHistory(data, filename)

		if data.Status == "DELIVERED" {
			fmt.Println("Update status")
			UpdateStatusDelivered(url)
		}
	}

	if isFileExist {
		isEqual := CompareUpdate(data, filename)

		if !isEqual {
			fmt.Println("Update resi")
			SendUpdateResi(data)
			SaveHistory(data, filename)

			if data.Status == "DELIVERED" {
				fmt.Println("Update status")
				UpdateStatusDelivered(url)
			}
		}
	}

}
