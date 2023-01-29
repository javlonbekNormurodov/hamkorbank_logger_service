package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"hamkorbank/config"
	"hamkorbank/events"
	"hamkorbank/pkg/logger"
)

func main() {

	cfg := config.Load()
	log := logger.NewLogger(cfg.LogLevel, logger.LevelInfo)

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	pubSubServer, err := events.NewEvents(cfg, log, ch)
	if err != nil {
		log.Error("error on event server")
	}

	ctx := context.Background()
	pubSubServer.InitServices(ctx) // it should run forever if there is any consumer
}
