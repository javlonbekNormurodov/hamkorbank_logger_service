package events

import (
	"context"
	"hamkorbank/config"
	"hamkorbank/events/logger_service"
	"hamkorbank/pkg/logger"
	"hamkorbank/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PubSubServer struct {
	cfg      config.Config
	rabbitmq rabbitmq.RabbitMQI
	log      logger.LoggerI
}

func NewEvents(cfg config.Config, log logger.LoggerI, ch *amqp.Channel) (*PubSubServer, error) {
	rabbit, err := rabbitmq.NewRabbitMQ(cfg, ch)
	if err != nil {
		return nil, err
	}

	return &PubSubServer{
		cfg:      cfg,
		log:      log,
		rabbitmq: rabbit,
	}, nil
}

func (s *PubSubServer) InitServices(ctx context.Context) {
	triggerListenerService := logger_service.NewTriggerListenerService(s.log, s.rabbitmq)
	triggerListenerService.RegisterConsumers()
	s.rabbitmq.RunConsumers(ctx)
}
