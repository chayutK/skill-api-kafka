package producer

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "net/http/pprof"

	"github.com/rcrowley/go-metrics"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	broker    = os.Getenv("BROKER")
	version   = sarama.DefaultVersion.String()
	topic     = os.Getenv("TOPIC")
	producers = 1
	verbose   = false

	recordsNumber int64 = 1
	recordsRate         = metrics.GetOrRegisterMeter("records.rate", nil)
)

func init() {
	fmt.Println(broker)
	if len(topic) == 0 {
		println(topic)
		panic("no topic given to be consumed, please set the -topic flag")
	}
}

func SendMessage(method string, message []byte) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(strings.Split(broker, ","), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	start := time.Now()
	msg := &sarama.ProducerMessage{Topic: "my-topic", Key: sarama.StringEncoder(method), Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
	fmt.Println("Time Usage :", time.Since(start))

}
