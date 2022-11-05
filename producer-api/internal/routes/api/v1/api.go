package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/config"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/controllers"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	rabbitMQBroker := config.NewRabbitMQBroker()
	apiGroup := router.Group(prefix)
	{
		apiV1 := apiGroup.Group("/rabbit-mq")
		{
			rabbitService := services.NewRabbitMQService(rabbitMQBroker)
			rabbitController := controllers.NewRabbitMQContoller(rabbitService)

			// api to send message into queue for RabbitMQ
			apiV1.POST("/push", rabbitController.SendMessages)
		}
	}
}
