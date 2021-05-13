package rabbit

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/streadway/amqp"

	"github.com/moooll/company-manager/rabbit/domain"
)

const MQURL = "MQ_URL"

type Msg struct {
	Function string
	Entity   domain.Entity
	Error    error
}

func Send(infoToSend Msg, queue string, ch amqp.Channel) error {
	body, err := ffjson.Marshal(infoToSend)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func Receive(queue string, k chan (Msg), ch amqp.Channel) (err error) {
	// conn, err := amqp.Dial(MQURL)
	// if err != nil {
	// 	return err
	// }

	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	return err
	// }

	// defer ch.Close()

	// q, err := amqp.QueueDeclare(
	// 	queue,
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	msgs, err := ch.Consume(
		queue,
		"",    // consumer
		false, // auto-ack false
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		return err
	}
	wait := make(chan bool)
	er := make(chan error)
	var mes Msg
	go func() {
		for m := range msgs {
			err = ffjson.Unmarshal(m.Body, &mes)
			if err != nil {
				er <- err
			}

			k <- mes
			m.Ack(true)
		}
	}()

	<-wait
	return nil
}
