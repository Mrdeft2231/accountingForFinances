package app

import (
	"accounting/pkg/logger"
	"accounting/services/gatway_accuntingForFinances/config"
	"accounting/services/gatway_accuntingForFinances/postgres"
	"context"
)

func InitApp() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	lg := logger.InitLg()

	_, err = postgres.InitDB(context.Background(), cfg.Postgres, lg)
	if err != nil {
		panic(err)
	}

}
