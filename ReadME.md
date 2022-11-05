#Golang rabbit-mq example


## How to run rabbitmq and application
- docker-compose build --no-cache
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f {service-name}
- send POST request to API' , http://127.0.0.1:8080/api/v1/rabbit-mq/push , for pushing messages
- docker-compose down (if you want to stop and remove containers)

____

```
Body
{
    "id":1,
    "fio":"Sanzhar Anarbay",
    "department":"Test",
    "age":23,
    "mark":3.45
}
```

____

- visit http://localhost:15672/ - (GUI RabbitMQ management ) , (login, password shown in .env file)
- visit http://localhost:9000/ - Portainer (GUI Docker containers)
- In portainer open producer-api container and see the logs , to check the message delivery status
- In portainer open consumer and see the logs , to check is consumer handle the message from the queue
- There are in general 3 .env files, in root directory , also in the consumer and producer-api directories
____
