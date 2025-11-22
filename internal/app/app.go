package app

import (
	"accounting/config"
	"accounting/internal/commands"
	"accounting/pkg/Postgers"
	"accounting/pkg/logger"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

func Message(ctx context.Context, b *bot.Bot, update *models.Update) {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	lg := logger.InitLg()

	_, err = Postgers.InitDB(context.Background(), cfg.Postgres, lg)
	if err != nil {
		panic(err)
	}
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
