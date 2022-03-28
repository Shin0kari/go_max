package main

import (
	"log"
	"os"

	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/handler"
	rep "github.com/Shin0kari/go_max/package/repository"
	sv "github.com/Shin0kari/go_max/package/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := rep.NewPostgresDB(rep.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// добавляем указатели на сервисы
	rep := rep.NewRepository(db)
	// конструктор для внедрения зависимостей сервиса
	services := sv.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(serv.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
