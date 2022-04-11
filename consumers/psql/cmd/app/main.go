package main

import (
	"log"
	"psql/database"
	"psql/rabbit"
)

func main() {
	psql, closeDB, err := database.New()
	if err != nil {
		log.Println("New database: ", err)
		return
	}
	defer closeDB()

	r, err := rabbit.Start()
	if err != nil {
		log.Println("connecting to rabbitmq: ", err)
		return
	}

	log.Println(r.Listen(psql))
}
