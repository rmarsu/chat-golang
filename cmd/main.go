package main

import (
	chat "chat"
	"chat/pkg/handler"
	"chat/pkg/repository"
	"chat/pkg/service"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	dbConnect, err := repository.NewDatabase()
	if err != nil {
		log.Fatal("database dead")
	}
	userRep := repository.NewRepository(dbConnect.GetDB())
	userSvc := service.NewService(userRep)
	hub := handler.NewHub()
	userHand := handler.NewHandler(userSvc, hub)
	go hub.Run()

	srv := new(chat.Server)
	if err := srv.Run("8080", userHand.InitRoutes()); err != nil {
		logrus.Fatal("runerr", err.Error())
	}
}
