package main

import (
	"log"

	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/handler"
	"github.com/Shin0kari/go_max/package/repository"
	"github.com/Shin0kari/go_max/package/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "9865guide",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// добавляем указатели на сервисы
	rep := repository.NewRepository(db)
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
