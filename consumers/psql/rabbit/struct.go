package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"psql/database"
)

type Rabbit struct {
	channel *amqp.Channel
	query   string
}

func (r Rabbit) Close() error {
	return r.channel.Close()
}

func (r Rabbit) Listen(psql *database.Postgres) error {

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

	log.Println("Listening Started")

	for m := range msg {
		log.Println("Update Started")

		var currencies = map[string]string{}
		err = json.Unmarshal(m.Body, &currencies)
		if err != nil {
			log.Println(fmt.Errorf("rabbitmq: unmarshal to map: %w", err))
			continue
		}

		err = psql.Insert(currencies)
		if err != nil {
			log.Println(fmt.Errorf("postgres: insert: %w", err))
			continue
		}
		log.Println("Updated")
	}

	return nil
}
