package main

import (
	"fmt"
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/spaceweasel/promptui"
	"log"
	"os"
	"piso-scrapper/scrapers"
	"time"
)

type fn func(string, *tgbotapi.BotAPI)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func runScraper(scraper []fn, url []string, bot *tgbotapi.BotAPI) {
	for {
		for exec := range scraper {
			go scraper[exec](url[exec], bot)
		}
		time.Sleep(15 * time.Minute)
	}

}

func main() {
	var selected = true
	var result []int
	var fns []fn
	var urls []string
	botToken := goDotEnvVariable("TELEGRAM_BOT_TOKEN")
	IdealistaURL := goDotEnvVariable("IDEALISTA_URL")
	FotocasaURL := goDotEnvVariable("FOTOCASA_URL")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("An error occurs")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for selected {
		prompt := promptui.MultiSelect{
			Label: "Select Web To scrap",
			Items: []string{"Idealista", "Fotocasa"},
		}
		result, err = prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if len(result) != 0 {
			selected = false
			break
		}
		fmt.Println("Please Select at least one option")
	}

	if len(result) == 1 {
		if result[0] == 1 {
			fns = append(fns, scrapers.ScrapF)
			if FotocasaURL == "" {
				log.Fatal("Introduce a valid URL")
			}
			urls = append(urls, FotocasaURL)
		} else {
			fns = append(fns, scrapers.ScrapI)
			if IdealistaURL == "" {
				log.Fatal("Introduce a valid URL")
			}
			urls = append(urls, IdealistaURL)
		}
	} else {
		fns = append(fns, scrapers.ScrapF)
		fns = append(fns, scrapers.ScrapI)
		if IdealistaURL == "" || FotocasaURL == "" {
			log.Fatal("Introduce a valid URL")
		}
		urls = append(urls, FotocasaURL)
		urls = append(urls, IdealistaURL)

	}

	runScraper(fns, urls, bot)

}
