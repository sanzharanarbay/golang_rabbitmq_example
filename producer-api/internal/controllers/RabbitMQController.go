package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/models"
	"github.com/sanzharanarbay/golang_rabbitmq_example/producer-api/internal/services"
	"log"
	"net/http"
)

type RabbitMQController struct {
	RabbitMQService *services.RabbitMQService
}

func NewRabbitMQContoller(RabbitMQService *services.RabbitMQService) *RabbitMQController {
	return &RabbitMQController{
		RabbitMQService: RabbitMQService,
	}
}

func (r *RabbitMQController) SendMessages(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	response := map[string]interface{}{}

	status, err := r.RabbitMQService.ProduceMessage(&user)
	if status == 1 {
		log.Println("Message successfully sent")
		response["status"] = true
		response["message"] = "Message successfully sent"
	} else {
		log.Println("Message send error")
		response["status"] = false
		response["message"] = err.Error()
	}
	c.JSON(http.StatusOK, response)
}