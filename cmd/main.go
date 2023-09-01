package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/rogudator/todo-app/internal/api"
	"github.com/rogudator/todo-app/internal/repository"
	"github.com/rogudator/todo-app/internal/service"
	"github.com/sirupsen/logrus"
)

//@title Todo App API
//@version 1.0
//@description API Server fro TodoList Application

//@host localhost:8080
//@basePath /

//@securityDefinition.apikey ApiKeyAuth
//@in header
//@name Authorization

func main() {
	logrus.Print("Begin")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variavles: %s", err.Error())
	}

	db, err := repository.NewPosgresDB(repository.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	})
	fmt.Println(repository.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := api.NewHandler(services)

	srv := new(api.Server)
	go func() {if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	}()

	logrus.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("TodoApp shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db disconnection: %s", err.Error())
	}
}