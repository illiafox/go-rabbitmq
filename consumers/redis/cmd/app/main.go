package main

import (
	"fmt"
	"log"

	"consumer/database"
	"consumer/rabbit"
	"consumer/server"
)

func main() {
	redis, closeDB, err := database.New()
	if err != nil {
		log.Println("Creating database: ", err)
		return
	}
	defer closeDB()

	r, err := rabbit.Start()
	if err != nil {
		log.Println("connecting to rabbitmq: ", err)
		return
	}
	defer r.Close()

	go func() {
		log.Println(server.Start(redis))
	}()

	fmt.Println(r.Listen(redis))
}
