package crawler

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/GnuYtuce/ytuyemekhane-api/models"
	"github.com/PuerkitoBio/goquery"
)

// YtuYemekhaneURL : yildiz teknik yemekhane listersini cekilecegi site.
var YtuYemekhaneURL = "http://www.sks.yildiz.edu.tr/yemekmenu/"

// Crawl : sitedeki menu datalarini cikartir.
func Crawl(date models.Date) (models.Menus, error) {
	var menus models.Menus
	if err := date.IsValid(); err != nil {
		return menus, err
	}

	url := createURL(date.Month, date.Year)
	log.Println("Crawl: ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []models.Menu{}, err
	}

	doc.Find("#menu_container").Each(func(i int, s *goquery.Selection) {
		lauchMenuStr := s.Find(".one_lunchMainMenu").Text()
		dinnerMenuStr := s.Find(".one_dinnerMainMenu").Text()
		dateStr := s.Find("#menuFooterFilter").Text()
		lauchMenuStr = clearText(lauchMenuStr)
		dinnerMenuStr = clearText(dinnerMenuStr)
		dateStr = clearText(dateStr)

		var menu = models.Menu{
			Lunch:  strings.Split(lauchMenuStr, "\n"),
			Dinner: strings.Split(dinnerMenuStr, "\n"),
			Date:   models.NewDate(dateStr),
		}

		menus = append(menus, menu)
	})

	if len(menus) > 0 {
		return menus, nil
	}

	return menus, errors.New("Bu tarihte hic menu yok")
}

func createURL(values ...interface{}) string {
	var result = YtuYemekhaneURL
	for _, value := range values {
		switch value.(type) {
		case int:
			result += intToString(value.(int))
		case string:
			result += value.(string)
		}
		result += "/"
	}
	return result
}

func clearText(orginal string) string {
	return strings.Trim(orginal, "\n\n ")
}

func intToString(value int) string {
	return strconv.Itoa(value)
}
