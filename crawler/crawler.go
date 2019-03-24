package crawler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/GnuYtuce/ytuyemekhane-api/models"
	"github.com/GnuYtuce/ytuyemekhane-api/util"
	"github.com/PuerkitoBio/goquery"
)

// Crawl : sitedeki menu datalarini cikartir.
func Crawl(url string) (models.Menus, error) {
	var menus models.Menus
	fmt.Printf("Crawl: %s\n", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []models.Menu{}, err
	}

	doc.Find("#menu_container").Each(func(i int, s *goquery.Selection) {
		lauchMenuStr := s.Find(".one_lunchMainMenu").Text()
		dinnerMenuStr := s.Find(".one_dinnerMainMenu").Text()
		dateStr := s.Find("#menuFooterFilter").Text()
		lauchMenuStr = util.ClearText(lauchMenuStr)
		dinnerMenuStr = util.ClearText(dinnerMenuStr)
		dateStr = util.ClearText(dateStr)

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
