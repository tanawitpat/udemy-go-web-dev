package main

import (
	"html/template"
	"log"
	"os"
)

type Hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		Hotel{
			Name:    "HotelA",
			Address: "100/100",
			City:    "Bangkok",
			Zip:     "10100",
			Region:  "Central",
		},
		Hotel{
			Name:    "HotelB",
			Address: "200/200",
			City:    "Chiang Mai",
			Zip:     "20100",
			Region:  "Northern",
		},
		Hotel{
			Name:    "HotelC",
			Address: "300/300",
			City:    "Phuket",
			Zip:     "30100",
			Region:  "Southern",
		},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
