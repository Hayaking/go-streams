package main

import (
	"context"
	"log"
	"strings"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/reugn/go-streams/flow"
	ext "github.com/reugn/go-streams/pulsar"
)

func main() {
	clientOptions := pulsar.ClientOptions{URL: "pulsar://localhost:6650"}
	producerOptions := pulsar.ProducerOptions{Topic: "test2"}
	consumerOptions := pulsar.ConsumerOptions{
		Topic:            "test1",
		SubscriptionName: "group1",
		Type:             pulsar.Exclusive,
	}

	ctx := context.Background()
	source, err := ext.NewPulsarSource(ctx, &clientOptions, &consumerOptions)
	if err != nil {
		log.Fatal(err)
	}
	flow1 := flow.NewMap(toUpper, 1)
	sink, err := ext.NewPulsarSink(ctx, &clientOptions, &producerOptions)
	if err != nil {
		log.Fatal(err)
	}

	source.
		Via(flow1).
		To(sink)
}

var toUpper = func(in interface{}) interface{} {
	msg := in.(pulsar.Message)
	return strings.ToUpper(string(msg.Payload()))
}
