package models

import (
	"fmt"
	"strings"

	"github.com/GnuYtuce/ytuyemekhane-api/util"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

const (
	DAY_START   = 1
	DAY_END     = 31
	MONTH_START = 1
	MONTH_END   = 12
	YEAR_START  = 2011
	YEAR_END    = 2030
)

func NewDate(dateStr string) *Date {
	var d Date
	temp := strings.Split(dateStr, ".")
	d.Day, _ = util.StringToInt(temp[0])
	d.Month, _ = util.StringToInt(strings.TrimLeft(temp[1], "0"))
	d.Year, _ = util.StringToInt(temp[2])
	return &d
}

func (d Date) ToString() string {
	return fmt.Sprintf("%d.%d.%d", d.Day, d.Month, d.Year)
}

func (d Date) IsValid() error {
	if err := d.isYearSuitable(); err != nil {
		return err
	}
	if err := d.isMonthSuitable(); err != nil {
		if d.Month != 0 {
			return err
		}
	}
	if err := d.isDaySuitable(); err != nil {
		if d.Day != 0 {
			return err
		}
	}
	return nil
}

func (d Date) isYearSuitable() error {
	if isInRange(d.Year, YEAR_START, YEAR_END) {
		return nil
	}
	return fmt.Errorf("%d ile %d arasinda bir yil secin:)", YEAR_START, YEAR_END)
}

func (d Date) isMonthSuitable() error {
	if isInRange(d.Month, MONTH_START, MONTH_END) {
		return nil
	}
	return fmt.Errorf("%d ile %d arasinda bir ay secin:)", MONTH_START, MONTH_END)
}

func (d Date) isDaySuitable() error {
	if isInRange(d.Day, DAY_START, DAY_END) {
		return nil
	}
	return fmt.Errorf("%d ile %d arasinda bir gun secin:)", DAY_START, DAY_END)
}

func isInRange(value, start, end int) bool {
	if value >= start && value <= end {
		return true
	}

	return false
}
