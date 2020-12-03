package cmd

import (
	"kafka-client/kafka"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	topic             string
	numMsg, batchSize int
)

// sendMsgCmd represents the sendMsg command
var sendMsgCmd = &cobra.Command{
	Use:   "sendMsg",
	Short: "Send messages to the kafka topic",
	Long:  `Send given numbers of messages with given batch-size to the kakfa topic`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
		err := kafka.SendMsg(kafkaURL, topic, numMsg, batchSize)
		if err != nil {
			log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(sendMsgCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendMsgCmd.PersistentFlags().String("foo", "", "A help for foo")
	sendMsgCmd.PersistentFlags().StringVar(&topic, "topic", "test", "The kafka topic to send messages to.")
	sendMsgCmd.PersistentFlags().IntVar(&numMsg, "num-msg", 1000, "Number of messages to send.")
	sendMsgCmd.PersistentFlags().IntVar(&batchSize, "batch-size", 1000, "Size of batch to send.")
}
