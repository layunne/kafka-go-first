package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {

	kafkaHost := "localhost:9092"
	topic := "topic-1"
	messages := []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaHost})

	if err != nil {
		panic(fmt.Sprintf("producer connect error: %v\n", err.Error()))
	}

	defer p.Close()

	for _, message := range messages {

		// Optional
		// Delivery report chan for produced message
		deliveredChan := make(chan kafka.Event)

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
			Headers:        []kafka.Header{{Key: "key-1", Value: []byte("value-1")}}, //Optional
		}

		//send message to kafka server
		err := p.Produce(message, deliveredChan)

		if err != nil {
			panic(fmt.Sprintf("send message error: %v\n", err.Error()))
		}

		// Optional
		go handlerDeliveredMessage(deliveredChan)

		// Wait to send a next message
		time.Sleep(time.Second)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func handlerDeliveredMessage(deliveredChan chan kafka.Event) {
	e := <- deliveredChan
	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
		} else {
			fmt.Printf("Delivered message: %v\n", string(ev.Value))
		}
	}
}
