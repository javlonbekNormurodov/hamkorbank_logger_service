package rabbitmq

import (
	"context"
	"errors"

	"github.com/rabbitmq/amqp091-go"
)

type HandlerFunc func(cancelFunc context.Context, data []byte) error

type Consumer struct {
	name    string
	handler HandlerFunc
	message <-chan amqp091.Delivery
}

func (r RabbitMQ) AddConsumer(name string, handler HandlerFunc) error {
	if _, ok := r.consumers[name]; ok {
		panic(errors.New("consumer with the same name already exists: " + name))
	}

	err := r.channel.ExchangeDeclare(
		"logs",  // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		panic(err)
	}

	q, err := r.channel.QueueDeclare(
		"queue."+name, // name
		false,         // durable
		false,         // delete when unused
		true,          // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	err = r.channel.QueueBind(
		q.Name, // queue name
		name,   // routing key
		"logs", // exchange
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	if err != nil {
		panic(err)
	}

	r.consumers[name] = &Consumer{
		name:    name,
		handler: handler,
		message: msgs,
	}

	return nil
}
