package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	/*
		now := time.Now()
		date := models.Date{
			Day:   now.Day(),
			Month: int(now.Month()),
			Year:  now.Year(),
		}
	*/
	date := models.Date{
		Day:   18,
		Month: 1,
		Year:  2018,
	}
	fmt.Println(date)
	url := YtuYemekhaneURL + util.CreateVirtualPATH(date.Month, date.Year)
	menus, _ := crawler.Crawl(url)
	for _, menu := range menus {
		fmt.Println("##########")
		menu.Print()
		fmt.Println(date.ToString())
		if menu.Date.ToString() == date.ToString() {
			sendJSON(w, menu)
			return
		}
	}
	sendJSON(w, models.Menu{})
}

// FoodListByCertainYear : belli bir yila gore yemek listesi donulecek.
func FoodListByCertainYear(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("FoodListByCertainYear function"))
}

// FoodListByCertainYearAndMonth : belli bir yil ve aya gore yemek listesi donulecek.
func FoodListByCertainYearAndMonth(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("FoodListByCertainYearAndMonth function"))
}

// FoodListByCertainTime : belli bir zamana gore yemek listesi donulecek.
func FoodListByCertainTime(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("FoodListByCertainTime function"))
}

// sendJSON :
func sendJSON(w http.ResponseWriter, data models.Menu) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	jContent, _ := json.MarshalIndent(data, "", " ")
	w.Write(jContent)
}
