package server

import (
	"log"
	"net/http"
	"os"

	"consumer/database"
	"consumer/server/methods"
)

func Start(redis *database.Redis) error {

	var m = methods.Methods{
		Redis: redis,
	}

	http.HandleFunc("/get", m.Update)

	port := ":" + os.Getenv("API_PORT")

	log.Println("API started at localhost" + port)

	return http.ListenAndServe(port, nil)
}
