package rabbitmq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func New(queue string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	))
	if err != nil {
		return nil, fmt.Errorf("dialing rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("opening channel: %w", err)

	}

	err = ch.ExchangeDeclare("currs.fanout", "fanout", false, true, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("declaring exchange: %w", err)
	}

	_, err = ch.QueueDeclare(
		queue,
		false,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("declaring queue: %w", err)
	}

	return &RabbitMQ{
		channel: ch,
	}, nil

}
