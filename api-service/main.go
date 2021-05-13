package main

import (
	"fmt"

	"go.uber.org/zap"
	"github.com/streadway/amqp"

	"api-service/cmd"
	"api-service/internal/messaging"

)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't initialize zap logger: %v", err)
		return
	}

	defer logger.Sync()

	conn, err := amqp.Dial(messaging.MQURL)
	if err != nil {
		zap.Error(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		zap.Error(err)
		return
	}

	defer ch.Close()
	_, err = ch.QueueDeclare(
		"to-db",
		false,
		false,
		false,
		false,
		nil,
	)
	zap.L().Info("to-db queue is declared")
	cmd.Serve()
}
