package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang-kafka/user-service/dto"
)

const (
	KafkaServer  = "localhost:9092"
	KafkaTopic   = "orders-v1-topic"
	KafkaGroupId = "user-service"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id":          "user-service-2",
		"auto.offset.reset": kafka.offset,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := KafkaTopic
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var order dto.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received Order: %+v\n", order)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}