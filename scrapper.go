package main

import (
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	m "piso-scrapper/models"
	"time"
)

func sendMessage(apartment m.Apartment, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(-704301832, "")
	msg.ParseMode = "html"

	msg.Text = "<b>NARCOPISO DETECTED</b>\n" +
		"<b>Dirección: </b>" + apartment.Direction + "\n" + "<b>Precio: </b>" + apartment.Price + " €\n" +
		"<a href=\"" + apartment.Url + "\">enlace</a>"

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	botToken := goDotEnvVariable("TELEGRAM_BOT_TOKEN")
	IdealistaURL := goDotEnvVariable("IDEALISTA_URL")
	FotocasaURL := goDotEnvVariable("FOTOCASA_URL")

	bot, err := tgbotapi.NewBotAPI(botToken)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	if err != nil {
		panic(err)
	}
	for {
		go scrapF(FotocasaURL, bot)
		go scrapI(IdealistaURL, bot)
		time.Sleep(15 * time.Minute)
	}

}
