package scrapers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	m "piso-scrapper/models"
)

func SendMessage(apartment m.Apartment, bot *tgbotapi.BotAPI,chat_id int) {
	msg := tgbotapi.NewMessage(chat_id, "")
	msg.ParseMode = "html"

	msg.Text = "<b>NEW</b>\n" +
		"<b>Dirección: </b>" + apartment.Direction + "\n" + "<b>Precio: </b>" + apartment.Price + " €\n" +
		"<a href=\"" + apartment.Url + "\">enlace</a>"

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
