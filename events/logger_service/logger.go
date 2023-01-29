package logger_service

import (
	"hamkorbank/config"
	"hamkorbank/pkg/logger"
	"hamkorbank/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	RecordId string `json:"record_id"`
}

type triggerListener struct {
	log      logger.LoggerI
	rabbitmq rabbitmq.RabbitMQI
	conn     *amqp.Connection
}

type NotFound struct {
	NotFound string `json:"not_found"`
}

type Phone struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

func (t *triggerListener) RegisterConsumers() {
	_ = t.rabbitmq.AddConsumer(config.AllErrors, t.ListenErrors)
	_ = t.rabbitmq.AddConsumer(config.AllInfo, t.ListenInfo)
	_ = t.rabbitmq.AddConsumer(config.AllDebug, t.ListenDebug)
	_ = t.rabbitmq.AddConsumer(config.All, t.ListenAll)
}

func NewTriggerListenerService(log logger.LoggerI, rabbit rabbitmq.RabbitMQI) *triggerListener {
	return &triggerListener{
		log:      log,
		rabbitmq: rabbit,
	}
}
