package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"publisher/api"
	"publisher/rabbitmq"
)

func main() {
	queue := os.Getenv("RABBITMQ_QUEUE_NAME")

	key := os.Getenv("API_KEY")

	if key == "" {
		log.Println("'API_KEY' variable is empty")
		return
	}

	every, err := strconv.Atoi(os.Getenv("API_EVERY"))
	if err != nil {
		log.Println("'API_EVERY' variable has wrong format: %w", err)
		return
	}

	// //

	r, err := rabbitmq.New(queue)
	if err != nil {
		log.Println("connecting to rabbitmq: ", err)
		return
	}

	defer r.Close()

	for {
		log.Println("Update started")

		m, err := api.Parse(key)
		if err != nil {
			log.Println("API: ", err)
		}

		data, err := json.Marshal(m)
		if err != nil {
			log.Println("map marshalling: ", err)
		}

		err = r.Publish(queue, "application/json", data)
		if err != nil {
			log.Println("publish: ", err)
		}

		log.Println("Updated")
		time.Sleep(time.Second * time.Duration(every))
	}
}
