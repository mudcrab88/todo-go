package main

import (
	"log"

	"github.com/mudcrab88/todo"
	"github.com/mudcrab88/todo/pkg/handler"
	"github.com/mudcrab88/todo/pkg/repository"
	"github.com/mudcrab88/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
