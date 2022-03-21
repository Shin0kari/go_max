package main

import (
	"log"

	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/handler"
	"github.com/Shin0kari/go_max/package/repository"
	"github.com/Shin0kari/go_max/package/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// добавляем указатели на сервисы
	rep := repository.NewRepository()
	// конструктор для внедрения зависимостей сервиса
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(serv.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
