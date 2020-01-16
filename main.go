package main

import (
	"net/http"
	"./controller"
	"./domain"
	"./service"
	"time"
	"github.com/jakewright/muxinator"
)

func main() {
	config := domain.Config{}

	configService := service.ConfigService{
		Config: &config,
		Location: "config.yaml",
	}

	go configService.Watch(time.Second * 3)

	c := controller.Controller{
		Config: &config,
	}

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", http.HandlerFunc(c.ReadConfig))
	router.ListenAndServe(":80")

}
