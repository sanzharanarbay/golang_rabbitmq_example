package services

import (
	"encoding/json"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/config"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/models"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQService struct {
	RabbitMQBroker *config.RabbitMQBroker
}

type RabbitMQServiceInterface interface {
	ProduceMessage(user *models.User) (int64, error)
}

func NewRabbitMQService(RabbitMQBroker *config.RabbitMQBroker) *RabbitMQService {
	return &RabbitMQService{
		RabbitMQBroker: RabbitMQBroker,
	}
}

func (r *RabbitMQService) ProduceMessage(user *models.User) (int64, error) {
	jsonBody, _ := json.Marshal(user)
	err := r.RabbitMQBroker.Channel.Publish(
		r.RabbitMQBroker.RabbitmqExchange,
		r.RabbitMQBroker.RabbitmqKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		log.Println(err.Error())
		panic(err)
		return 0, err
	}
	return 1, nil
}
