package cmd

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/bmurase/codepix/application/kafka"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		deliveryChannel := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()

		kafka.Publish("eita rapaz", "test", producer, deliveryChannel)
		kafka.DeliveryReport(deliveryChannel)

		fmt.Println("message produced")
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
