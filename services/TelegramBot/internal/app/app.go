package app

import (
	"accounting/services/TelegramBot/internal/commands"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

func InitApp(ctx context.Context) {

	opts := []bot.Option{
		bot.WithDefaultHandler(Message),
	}
	b, err := bot.New("7789547644:AAHW_4jbfPtcFCvd5YIttqrM87dJyy8ZymY", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)

}

func Message(ctx context.Context, b *bot.Bot, update *models.Update) {

	m := update.Message.Text
	if strings.HasPrefix(m, "/") {
		commands.Commands(ctx, b, update)
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}

}
