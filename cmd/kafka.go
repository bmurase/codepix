package cmd

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/bmurase/codepix/application/kafka"
	"github.com/bmurase/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		deliveryChannel := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()

		kafka.Publish("hello kafka", "test", producer, deliveryChannel)
		go kafka.DeliveryReport(deliveryChannel)

		fmt.Println("message produced")

		database := db.ConnectDB(os.Getenv("env"))

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
