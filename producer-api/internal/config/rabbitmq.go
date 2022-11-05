package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type RabbitMQBroker struct {
	Channel          *amqp.Channel
	RabbitmqHost     string
	RabbitmqVhost    string
	RabbitmqLogin    string
	RabbitmqPassword string
	RabbitmqQueue    string
	RabbitmqExchange string
	RabbitmqKey      string
	RabbitmqConsumer string
}

func NewRabbitMQBroker() *RabbitMQBroker {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqVhost := os.Getenv("RABBITMQ_VHOST")
	rabbitmqLogin := os.Getenv("RABBITMQ_LOGIN")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmqQueue := os.Getenv("RABBITMQ_QUEUE")
	rabbitmqExchange := os.Getenv("RABBITMQ_EXCHANGE")
	rabbitmqKey := os.Getenv("RABBITMQ_ROUTE_KEY")
	rabbitmqConsumer := os.Getenv("RABBITMQ_CONSUMER")

	rabbitUri := fmt.Sprintf("amqp://%s:%s@%s%s", rabbitmqLogin, rabbitmqPassword, rabbitmqHost, rabbitmqVhost) //Build connection string

	conn, err := amqp.Dial(rabbitUri)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	if err == nil {
		fmt.Println("Successfully Connected to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	//defer ch.Close()

	_, err = ch.QueueDeclare(
		rabbitmqQueue, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	err = ch.ExchangeDeclare(
		rabbitmqExchange, // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)

	err = ch.QueueBind(
		rabbitmqQueue,    // queue name
		rabbitmqKey,      // routing key
		rabbitmqExchange, // exchange
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	return &RabbitMQBroker{
		Channel:          ch,
		RabbitmqHost:     rabbitmqHost,
		RabbitmqVhost:    rabbitmqVhost,
		RabbitmqLogin:    rabbitmqLogin,
		RabbitmqPassword: rabbitmqPassword,
		RabbitmqQueue:    rabbitmqQueue,
		RabbitmqExchange: rabbitmqExchange,
		RabbitmqKey:      rabbitmqKey,
		RabbitmqConsumer: rabbitmqConsumer,
	}
}
