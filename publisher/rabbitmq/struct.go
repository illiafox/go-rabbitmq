package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel *amqp.Channel
}

func (r RabbitMQ) Close() error {
	return r.channel.Close()
}

func (r RabbitMQ) Publish(queue, contentType string, body []byte) error {

	err := r.channel.Publish(
		"currs.fanout",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("publishing: %w", err)
	}

	return nil
}
