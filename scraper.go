package main

import (
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"piso-scrapper/scraper"
	"time"
)

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
		go scraper.ScrapF(FotocasaURL, bot)
		go scraper.ScrapI(IdealistaURL, bot)
		time.Sleep(15 * time.Minute)
	}

}
