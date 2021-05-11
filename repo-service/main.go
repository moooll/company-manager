package main

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't initialize zap logger: %v", err)
		return
	}

	defer logger.Sync()

	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "pass",
		Database: "manage_app_db",
	})

	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		zap.L().Error("could not launch the db:" + err.Error())
		return
	}
	


}
