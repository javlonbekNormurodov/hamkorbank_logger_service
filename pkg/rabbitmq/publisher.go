package rabbitmq

import (
	"errors"
)

type Publisher struct {
	name string
}

func (r RabbitMQ) AddPublisher(name string) error {
	if _, ok := r.publishers[name]; ok {
		panic(errors.New("consumer with the same name already exists: " + name))
	}

	r.publishers[name] = &Publisher{
		name: name,
	}

	return nil
}

func (r RabbitMQ) Publish(name string, data interface{}) error {
	//result, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}
	//
	//p := r.publishers[name]
	//
	//if p == nil {
	//	return errors.New("publisher with that topic doesn't exists: " + name)
	//}

	//msg := sarama.ProducerMessage{
	//	Topic: topic,
	//	Value: sarama.ByteEncoder(result),
	//}
	//_, _, err = p.client.SendMessage(&msg)

	//if err != nil {
	//	return err
	//}

	return nil
}
