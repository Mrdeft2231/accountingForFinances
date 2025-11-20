package commands

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Commands(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch update.Message.Text {
	case "/start":
		update.Message.Text = "Добро пожаловать в бота, это пока что разработка"
	case "/command":
		update.Message.Text = "Вот представленные команды"
	case "/info":
		update.Message.Text = "доступные команды: /start, /command"
	default:
		update.Message.Text = "Не правильная команда"
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
