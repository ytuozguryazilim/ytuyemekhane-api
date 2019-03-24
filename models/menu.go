package models

import (
	"log"
)

// Menu : veri yapisinda o gunun tarihi ve o gundeki oglen ve aksam yemek listesi tutulur.
type Menu struct {
	Date   *Date    `json:"date"`
	Lunch  []string `json:"lunch"`
	Dinner []string `json:"dinner"`
}

func (m Menu) Print() {
	log.Printf("Date: %s, Lunch: %s, Dinner: %s", m.Date.ToString(), m.Lunch, m.Dinner)
}
