package consumer

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/chayutK/skill-api-kafka/comsumer/router"
)

type consumerGroupHandler struct {
	id int

	router router.Router
}

func NewConsumerGroup(id int, router router.Router) *consumerGroupHandler {
	return &consumerGroupHandler{id: id, router: router}
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumer consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: key = %s, value = %s, timestamp = %v, topic = %s", string(message.Key), string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
			consumer.router.Route(string(message.Key), message.Value)
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			fmt.Println("Rebalancing....")
			return nil
		}
	}
}
