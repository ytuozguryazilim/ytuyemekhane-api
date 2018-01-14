package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	// YtuYemekhaneURL : yildiz teknik yemekhane listersini cekilecegi site.
	YtuYemekhaneURL = "http://www.sks.yildiz.edu.tr/yemekmenu/"
)

// FoodListByCurrentTime : suanki zamana gore yemek listesi donulecek.
func FoodListByCurrentTime(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("FoodListByCurrentTime function"))
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
