package scraper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	m "piso-scrapper/models"
)

func SendMessage(apartment m.Apartment, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(-704301832, "")
	msg.ParseMode = "html"

	msg.Text = "<b>NARCOPISO DETECTED</b>\n" +
		"<b>Dirección: </b>" + apartment.Direction + "\n" + "<b>Precio: </b>" + apartment.Price + " €\n" +
		"<a href=\"" + apartment.Url + "\">enlace</a>"

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
