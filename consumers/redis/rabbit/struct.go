package rabbit

import (
	"consumer/database"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Rabbit struct {
	channel *amqp.Channel
	query   string
}

func (r Rabbit) Close() error {
	return r.channel.Close()
}

func (r Rabbit) Listen(redis *database.Redis) error {

	msg, err := r.channel.Consume(
		r.query,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("consuming channel: %w", err)
	}

	log.Println("Listening Started ")

	for m := range msg {
		var currencies = map[string]string{}
		err = json.Unmarshal(m.Body, &currencies)
		if err != nil {
			log.Println(fmt.Errorf("rabbitmq: unmarshal to map: %w", err))
			continue
		}

		for k, v := range currencies {
			err = redis.Currencies.Set(k, v)
			if err != nil {
				log.Println(fmt.Errorf("rabbitmq: set currency: %w", err))
			}
		}
		log.Println("Updated")
	}

	return nil
}
