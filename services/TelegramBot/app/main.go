package main

import (
	"accounting/services/TelegramBot/internal/app"
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app.InitApp(ctx)

}
