package models

import (
	"strings"

	"github.com/GnuYtuce/ytuyemekhane-api/util"
)

// Date :
type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// ToString : Date veri yapisinin elemanlari "string" olarak doner. Misal "1.12.1999"
func (d Date) ToString() string {
	return strings.Join([]string{util.IntToString(d.Day), util.IntToString(d.Month), util.IntToString(d.Year)}, ".")
}

// Set : Date veri yapisinin elemanlari ayarlanir.
// Misal "Set" fonksiyonuna "10.01.2005" seklinde "string" gecilir. "Set" fonksiyonu da bunu parse ederek, kendi elemanlarina atar.
func (d *Date) Set(dateStr string) {
	temp := strings.Split(dateStr, ".")
	d.Day, _ = util.StringToInt(temp[0])
	d.Month, _ = util.StringToInt(strings.TrimLeft(temp[1], "0"))
	d.Year, _ = util.StringToInt(temp[2])
}
