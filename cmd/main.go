package main

import (
	"log"

	"github.com/Shin0kari/go_max/package/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(serv.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

//1
