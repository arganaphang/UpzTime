package main

import (
	"application/pkg/sendz/telegram"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
)

func main() {
	interval := os.Getenv("INTERVAL")
	if interval == "" {
		interval = "5s"
	}
	endpoint, ok := os.LookupEnv("ENDPOINT")
	if !ok {
		log.Fatal("please set ENDPOINT variable")
	}
	token, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok || token == "" {
		log.Fatal("please set TELEGRAM_TOKEN variable")
	}
	chatID, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		log.Fatal("please set TELEGRAM_CHAT_ID variable")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := resty.New()
	s := gocron.NewScheduler(time.UTC)
	t, err := telegram.New(token, int64(chatID))
	if err != nil {
		log.Fatal("can't connect into telegram bot")
	}
	last := true
	s.Every(interval).Do(func() {
		res, _ := c.R().Get(endpoint)
		// failed
		if res.StatusCode() != http.StatusOK && last {
			log.Printf("%s service is DOWN\n", endpoint)
			last = false
			t.Send(ctx, fmt.Sprintf("%s service DOWN", endpoint))
			return
		}
		// success
		if res.StatusCode() == http.StatusOK && !last {
			log.Printf("%s service is UP\n", endpoint)
			last = true
			t.Send(ctx, fmt.Sprintf("%s service UP", endpoint))
		}
	})

	s.StartBlocking()
}
