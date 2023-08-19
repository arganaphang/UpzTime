package telegram

import (
	"context"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Provider struct {
	bot    *api.BotAPI
	chatID int64
}

func New(token string, chatID int64) (*Provider, error) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = false
	return &Provider{
		bot:    bot,
		chatID: chatID,
	}, nil
}

func (p Provider) Send(ctx context.Context, msg string) error {
	_, err := p.bot.Send(api.NewMessage(p.chatID, msg))
	return err
}
