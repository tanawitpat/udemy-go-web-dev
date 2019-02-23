package main

import (
	"html/template"
	"log"
	"os"
)

type Restaurant struct {
	Name  string
	Menus []Menu
}
type Menu struct {
	Name     string
	Price    float64
	Category string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []Restaurant{
		Restaurant{
			Name: "Hong Kong House",
			Menus: []Menu{
				{
					Name:     "เป็ดทรงเครื่อง",
					Price:    500,
					Category: "Breakfast",
				},
				{
					Name:     "ปลากระพงทอดน้ำปลา",
					Price:    390,
					Category: "Lunch",
				},
			},
		},
		Restaurant{
			Name: "Hua Seng Hong",
			Menus: []Menu{
				{
					Name:     "ผัดคะน้า",
					Price:    200,
					Category: "Breakfast",
				},
				{
					Name:     "เป็ดย่าง",
					Price:    600,
					Category: "Dinner",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
