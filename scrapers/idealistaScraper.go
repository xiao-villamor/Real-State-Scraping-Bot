package scrapers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly"
	m "piso-scrapper/models"
	apartmentService "piso-scrapper/services/apartment.service"
	"time"
)

func ScrapI(url string, bot *tgbotapi.BotAPI) {
	fmt.Println(url)

	var err error
	iFactory, _ := apartmentService.GetApartmentFactory("idealista")

	c := colly.NewCollector(
		colly.AllowedDomains("idealista.com", "www.idealista.com"),
		colly.MaxBodySize(0),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0"),
	)

	detailcolector := c.Clone()

	c.OnHTML("a[class=item-link]", func(e *colly.HTMLElement) {
		url := "https://www.idealista.com" + e.Attr("href")
		detailcolector.Visit(url)

	})

	detailcolector.OnHTML("div[class=container]", func(e *colly.HTMLElement) {

		if e.ChildText(".main-info__title-main") != "" {

			a := m.Apartment{
				Id:        e.ChildText(".txt-ref"),
				Direction: e.ChildText(".main-info__title-minor"),
				Price:     e.ChildText("span.txt-bold"),
				Url:       e.Request.URL.String(),
			}

			if iFactory.FindCoincidences(a) {
				time.Sleep(100 * time.Millisecond)
				SendMessage(a, bot)

				iFactory.Create(a)
				time.Sleep(4 * time.Second)
			}
		}
	})

	err = c.Visit(url)
	if err != nil {
		return
	}

}
