package rabbit

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func Start() (*Rabbit, error) {
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

	q, err := ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when usused
		true,  //exclusive
		false, //no-wait
		nil,   //arguments
	)
	if err != nil {
		return nil, fmt.Errorf("declaring queue: %w", err)
	}

	err = ch.QueueBind(
		q.Name,         //queue name
		"",             //routing key
		"currs.fanout", //exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("binding queue: %w", err)
	}

	return &Rabbit{
		channel: ch,
	}, nil

}
