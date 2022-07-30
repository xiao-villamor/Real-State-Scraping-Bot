package scrapers

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"piso-scrapper/models"
	apartmentService "piso-scrapper/services/apartment.service"
	"strconv"
	"time"
)

func ScrapF(url string, bot *tgbotapi.BotAPI) {

	fmt.Println("starting scrapper")

	var err error
	var resp *http.Response
	iFactory, _ := apartmentService.GetApartmentFactory("fotocasa")
	err = iFactory.UpdateCount(12)
	if err != nil {
		return
	}
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string][]map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	for _, s := range result["realEstates"] {

		var url = s["detail"].(map[string]interface{})
		var id = s["id"]
		var directiontmp = s["address"].(map[string]interface{})
		var dir = directiontmp["ubication"].(string)
		var transactions = s["transactions"].([]interface{})
		var value = transactions[0].(map[string]interface{})
		var price = value["value"].([]interface{})
		var priceint = int(price[0].(float64))
		var priceString = strconv.Itoa(priceint)

		a := models.Apartment{
			Id:        fmt.Sprintf("%f", id.(float64)),
			Direction: dir,
			Price:     priceString,
			Url:       "https://www.fotocasa.es" + url["es"].(string),
		}
		if iFactory.FindCoincidences(a) {
			time.Sleep(100 * time.Millisecond)
			SendMessage(a, bot)
			iFactory.Create(a)
			time.Sleep(4 * time.Second)

		}

	}
}
