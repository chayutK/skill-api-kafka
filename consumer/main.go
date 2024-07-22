package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"

	sarama "github.com/IBM/sarama"
	consumer "github.com/chayutK/skill-api-kafka/comsumer/instance"
	"github.com/chayutK/skill-api-kafka/comsumer/repository/database"
	"github.com/chayutK/skill-api-kafka/comsumer/router"
)

func main() {
	config := sarama.NewConfig()
	// config.Version = sarama.V2_0_0_0 // specify appropriate version
	// config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	group, err := sarama.NewConsumerGroup([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, "my-group", config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	db := database.Sync()
	router := router.NewRouter(db)

	// Iterate over consumer sessions.
	// ctx := context.Background()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()
	for {
		topics := []string{"my-topic"}
		handler := consumer.NewConsumerGroup(1, router)

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := group.Consume(ctx, topics, handler); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return
			}
			panic(err)
		}
	}
}
