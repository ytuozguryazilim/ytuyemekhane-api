package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GnuYtuce/ytuyemekhane-api/crawler"
	"github.com/GnuYtuce/ytuyemekhane-api/models"
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
	url := YtuYemekhaneURL + util.CreateVirtualPATH(date.Month, date.Year)
	menus, _ := crawler.Crawl(url)
	menu, err := menus.SearchMenuByDate(date)
	if err != nil {
		sendJSON(w, models.Menu{})
		return
	}
	sendJSON(w, menu)
}

// FoodListByCertainYear : belli bir yila gore yemek listesi donulecek.
func FoodListByCertainYear(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("FoodListByCertainYear function"))
}

// FoodListByCertainYearAndMonth : belli bir yil ve aya gore yemek listesi donulecek.
func FoodListByCertainYearAndMonth(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	month, _ := util.StringToInt(p.ByName("month"))
	year, _ := util.StringToInt(p.ByName("year"))
	url := YtuYemekhaneURL + util.CreateVirtualPATH(month, year)
	menus, _ := crawler.Crawl(url)
	sendJSON(w, menus)
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
	url := YtuYemekhaneURL + util.CreateVirtualPATH(date.Month, date.Year)
	menus, _ := crawler.Crawl(url)
	menu, err := menus.SearchMenuByDate(date)
	if err != nil {
		sendJSON(w, models.Menu{})
		return
	}
	sendJSON(w, menu)
}

// sendJSON :
func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	jContent, _ := json.MarshalIndent(data, "", " ")
	w.Write(jContent)
}
