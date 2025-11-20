package main

import (
	"accounting/internal/app"
	"context"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(app.Message),
	}

	b, err := bot.New("7789547644:AAHW_4jbfPtcFCvd5YIttqrM87dJyy8ZymY", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
