package app

import (
	"accounting/internal/commands"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

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
