package main

import (
	"log"
	"os"
	"todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err:=initConfig(); err!=nil{
		log.Fatalf("Error with config file read")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: os.Getenv("DBHost"),
		Port: os.Getenv("DBPort"),
		Username: os.Getenv("DBUsername"),
		Password: os.Getenv("DBPassword"),
		DBName: os.Getenv("DBName"),
		SSLMode: os.Getenv("DBSSSLMode"),
	})

	if err != nil {
		log.Fatalf("Error with db connection: %s", err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)

	if err := srv.Run(os.Getenv("AppPort"), handler.InitRoutes()); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}

func initConfig() error {
	if err:=godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error with .env file loading. Error: %s", err)
		return err
	}
	return nil
}