package main

import (
	"html/template"
	"log"
	"os"
)

type Restaurant struct {
	Name  string
	Menus MenuCategory
}
type Menu struct {
	Name  string
	Price float64
}

type MenuCategory struct {
	Breakfast, Lunch, Dinner []Menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []Restaurant{
		Restaurant{
			Name: "Hong Kong House",
			Menus: MenuCategory{
				Breakfast: []Menu{
					{
						Name:  "ปลากระพงทอดน้ำปลา",
						Price: 390,
					},
					{
						Name:  "ผัดคะน้า",
						Price: 200,
					},
					{
						Name:  "เป็ดย่าง",
						Price: 600,
					},
				},
				Lunch: []Menu{
					{
						Name:  "ปูผัดผงกะหรี่",
						Price: 400,
					},
				},
				Dinner: []Menu{
					{
						Name:  "ขนมจีบหมู",
						Price: 100,
					},
				},
			},
		},
		Restaurant{
			Name: "Hua Seng Hong",
			Menus: MenuCategory{
				Breakfast: []Menu{
					{
						Name:  "ไข่เจียวหอยนางรม",
						Price: 100,
					},
					{
						Name:  "ออส่วน",
						Price: 120,
					},
				},
				Lunch: []Menu{
					{
						Name:  "กุ้งแม่น้ำ",
						Price: 300,
					},
				},
				Dinner: []Menu{
					{
						Name:  "กุ้งแช่น้ำปลา",
						Price: 200,
					},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
