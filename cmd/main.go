package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handler := new(handler.Handler)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}