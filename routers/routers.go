package routers

import (
	"errors"
	"net/http"
	"sync"
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
	if err := date.IsValid(); err != nil {
		sender.Err(w, err)
		return
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
	year, _ := util.StringToInt(p.ByName("year"))
	date := models.Date{
		Day:   0,
		Month: 0,
		Year:  year,
	}
	if err := date.IsValid(); err != nil {
		sender.Err(w, err)
		return
	}

	var AllMenus models.Menus
	c := make(chan models.Menus, 12)

	var wg sync.WaitGroup
	wg.Add(12)
	for month := 1; month <= 12; month++ {
		url := util.CreateURL(YtuYemekhaneURL, month, year)
		go func(url string) {
			defer wg.Done()
			menus, _ := crawler.Crawl(url)
			c <- menus
		}(url)
	}
	wg.Wait()
	close(c)
	for menus := range c {
		menus.Print()
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
	date := models.Date{
		Day:   0,
		Month: month,
		Year:  year,
	}
	if err := date.IsValid(); err != nil {
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
	date := models.Date{
		Day:   day,
		Month: month,
		Year:  year,
	}
	if err := date.IsValid(); err != nil {
		sender.Err(w, err)
		return
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
