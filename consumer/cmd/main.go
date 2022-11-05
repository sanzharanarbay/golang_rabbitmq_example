package main

import (
	"fmt"
	"github.com/streadway/amqp"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"log"
)

func main() {
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
		log.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	defer conn.Close()

	if err == nil {
		log.Println("Successfully Connected to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

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

	messages, err := ch.Consume(
		rabbitmqQueue,    // queue name
		rabbitmqConsumer, // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no local
		false,            // no wait
		nil,              // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-forever

}
