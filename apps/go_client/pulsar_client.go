package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var client pulsar.Client

func NewPulsarClient() {

	var err error
	pulsar_url := fmt.Sprintf("pulsar://%s:%s", C.Pulsar.Url, C.Pulsar.Port)

	log.Infow("NewPulsarClient started",
		"pulsar_url", pulsar_url,
	)

	client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsar_url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()
}

func SendMessageAsProducer(topic string, msg string) {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})

	defer producer.Close()

	if err != nil {
		log.Infof("Failed to publish message", err)
	}
	log.Infow("Published message",
		"topic", topic,
		"Message", msg,
	)
}

type messageSchema struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Consume(topic string) {

	options := pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "sub1",
		Type:             pulsar.Shared,
	}
	// create consumer
	consumer, err := client.Subscribe(options)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

}
