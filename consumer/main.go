package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("starting...")

	host := "localhost:9092"
	topic := "topic-1"
	group := "group-1"

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	// subscribe in topic
	err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

	defer c.Close()

	fmt.Println("pulling...")
	for {
		message, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("pulling error: %v\n", err.Error())
			return
		} else {
			fmt.Printf("message %s: %s \n", message.TopicPartition, string(message.Value))

			for _, header := range message.Headers {
				fmt.Printf("header %s - [%s:%s]\n", message.TopicPartition, header.Key, string(header.Value))
			}

		}
	}
}
