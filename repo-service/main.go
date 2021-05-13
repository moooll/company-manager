package main

import (
	"context"
	"fmt"
	"os"

	pg "github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"github.com/streadway/amqp"

	"github.com/moooll/company-manager/repo-service/internal/constants"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't initialize zap logger: %v", err)
		return
	}
	zap.ReplaceGlobals(logger)

	defer logger.Sync()

	conn, err := amqp.Dial(constants.MQURL)
	if err != nil {
		zap.Error(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		zap.Error(err)
	}

	defer ch.Close()
	_, err = ch.QueueDeclare(
		"from-db",
		false,
		false,
		false,
		false,
		nil,
	)
	zap.L().Info("to-db queue is declared")
	pseudoDsn := fmt.Sprintf(
		"%s:%s",
		os.Getenv(constants.DatabaseHost),
		os.Getenv(constants.DatabasePort),
	)

	db := pg.Connect(&pg.Options{
		Addr:     pseudoDsn,
		User:     os.Getenv(constants.DatabaseUser),
		Password: os.Getenv(constants.DatabasePassword),
		Database: os.Getenv(constants.DatabaseName),
	})

	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		println(err.Error())
		zap.L().Error("could not launch the db:" + err.Error())
		return
	}

	zap.L().Info("service is running")

	select {}
}
