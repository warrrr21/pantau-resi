package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ListResi struct {
	Link   string `json:"link"`
	Status string `json:"status"`
}

type ListResiDetail struct {
	Resi []ListResi `json:"resi"`
}

func ReadListResi() ListResiDetail {

	// read file resi.json
	file, err := ioutil.ReadFile("resi.json")
	if err != nil {
		log.Fatal(err)
	}

	// marshal json to struct
	var data ListResiDetail
	json.Unmarshal(file, &data)

	return data
}

func GetHtml(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

type History struct {
	Id          int    `json:"id"`
	Process     string `json:"process"`
	Description string `json:"description"`
}

type Resi struct {
	NoResi  string    `json:"resi"`
	Status  string    `json:"status"`
	History []History `json:"history"`
}

func ParseHtml(doc *goquery.Document) Resi {

	status := strings.TrimSpace(doc.Find("h4").First().Text())
	resi := strings.TrimSpace(doc.Find("h5").First().Text())

	r := Resi{Status: status, NoResi: resi}

	body := doc.Find(".list-custom.list-custom-circle")

	body.Find("li").Each(func(i int, s *goquery.Selection) {
		process := s.Find("h5").Text()

		desc := s.Find("span").Text()

		h := History{Id: i + 1, Process: process, Description: desc}

		r.History = append(r.History, h)
	})

	return r
}

func DoesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func CompareUpdate(resi Resi, filename string) bool {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var data Resi
	json.Unmarshal(file, &data)

	return reflect.DeepEqual(data, resi)
}

func SendUpdateResi(resi Resi) {

	body := fmt.Sprintf("[Update]\n\nNomor Resi:%s\n\nHistory:\n\n", resi.NoResi)

	for i := range resi.History {
		b := fmt.Sprintf("%d. %s\n%s", resi.History[i].Id, resi.History[i].Process, resi.History[i].Description)

		if i == len(resi.History)-1 {
			body = body + b
			continue
		}
		body = body + b + "\n\n"
	}

	SendMessage(body)
}

func SaveHistory(resi Resi, filename string) {
	file, _ := json.MarshalIndent(resi, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)

}

func UpdateStatusDelivered(url string) {
	listResi := ReadListResi()

	for i := range listResi.Resi {
		if listResi.Resi[i].Link == url {
			listResi.Resi[i].Status = "DELIVERED"
		}
	}

	fmt.Println("Update status")

	file, _ := json.MarshalIndent(listResi, "", " ")
	_ = ioutil.WriteFile("resi.json", file, 0644)
}
