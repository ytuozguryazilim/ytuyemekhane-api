package routers

import (
	"errors"
	"net/http"
	"time"

	"github.com/GnuYtuce/ytuyemekhane-api/crawler"
	"github.com/GnuYtuce/ytuyemekhane-api/models"
	"github.com/GnuYtuce/ytuyemekhane-api/sender"
	"github.com/GnuYtuce/ytuyemekhane-api/util"
	"github.com/julienschmidt/httprouter"
)

var (
	// YtuYemekhaneURL : yildiz teknik yemekhane listersini cekilecegi site.
	YtuYemekhaneURL = "http://www.sks.yildiz.edu.tr/yemekmenu"
)

// FoodListByCurrentTime : suanki zamana gore yemek listesi donulecek.
func FoodListByCurrentTime(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	now := time.Now()
	date := models.Date{
		Day:   now.Day(),
		Month: int(now.Month()),
		Year:  now.Year(),
	}
	url := util.CreateURL(YtuYemekhaneURL, date.Month, date.Year)
	menus, err := crawler.Crawl(url)
	if err != nil {
		sender.Err(w, err)
		return
	}
	menu, err := menus.SearchMenuByDate(date)
	if err != nil {
		sender.Err(w, err)
		return
	}
	sender.JSON(w, menu)
}

// FoodListByCertainYear : belli bir yila gore yemek listesi donulecek.
func FoodListByCertainYear(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var AllMenus models.Menus
	year, _ := util.StringToInt(p.ByName("year"))
	if err := util.IsYearSuitable(year); err != nil {
		sender.Err(w, err)
		return
	}
	for month := 1; month <= 12; month++ {
		url := util.CreateURL(YtuYemekhaneURL, month, year)
		menus, err := crawler.Crawl(url)
		if err != nil {
			continue
		}
		AllMenus = append(AllMenus, menus...)
	}
	if len(AllMenus) > 0 {
		sender.JSON(w, AllMenus)
		return
	}
	sender.Err(w, errors.New("Bu yila ait hic menu yok"))
}

// FoodListByCertainYearAndMonth : belli bir yil ve aya gore yemek listesi donulecek.
func FoodListByCertainYearAndMonth(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	month, _ := util.StringToInt(p.ByName("month"))
	year, _ := util.StringToInt(p.ByName("year"))
	if err := util.IsYearSuitable(year); err != nil {
		sender.Err(w, err)
		return
	}
	if err := util.IsMonthSuitable(month); err != nil {
		sender.Err(w, err)
		return
	}
	url := util.CreateURL(YtuYemekhaneURL, month, year)
	menus, err := crawler.Crawl(url)
	if err != nil {
		sender.Err(w, err)
		return
	}
	sender.JSON(w, menus)
}

// FoodListByCertainTime : belli bir zamana gore yemek listesi donulecek.
func FoodListByCertainTime(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	day, _ := util.StringToInt(p.ByName("day"))
	month, _ := util.StringToInt(p.ByName("month"))
	year, _ := util.StringToInt(p.ByName("year"))
	if err := util.IsYearSuitable(year); err != nil {
		sender.Err(w, err)
		return
	}
	if err := util.IsMonthSuitable(month); err != nil {
		sender.Err(w, err)
		return
	}
	if err := util.IsDaySuitable(day); err != nil {
		sender.Err(w, err)
		return
	}
	date := models.Date{
		Day:   day,
		Month: month,
		Year:  year,
	}
	url := util.CreateURL(YtuYemekhaneURL, date.Month, date.Year)
	menus, err := crawler.Crawl(url)
	if err != nil {
		sender.Err(w, err)
		return
	}
	menu, err := menus.SearchMenuByDate(date)
	if err != nil {
		sender.Err(w, err)
		return
	}
	sender.JSON(w, menu)
}
