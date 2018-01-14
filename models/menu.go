package models

import "fmt"

// Menu : veri yapisinda o gunun tarihi ve o gundeki oglen ve aksam yemek listesi tutulur.
type Menu struct {
	Date   *Date    `json:"date"`
	Lunch  []string `json:"lunch"`
	Dinner []string `json:"dinner"`
}

// Print : Menu veri yapisindaki bilgileri ekrana basar.
func (m Menu) Print() {
	fmt.Println("Date: ", m.Date.ToString())
	fmt.Println("Lunch:")
	for _, l := range m.Lunch {
		fmt.Println(l)
	}
	fmt.Println("Dinner:")
	for _, d := range m.Dinner {
		fmt.Println(d)
	}
}
