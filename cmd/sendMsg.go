package cmd

import (
	"kafka-client/kafka"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	topic                           string
	numMsg, targetOffset, batchSize int64
)

// sendMsgCmd represents the sendMsg command
var sendMsgCmd = &cobra.Command{
	Use:   "sendMsg",
	Short: "Send messages to the kafka topic",
	Long:  `Send given numbers of messages with given batch-size to the kakfa topic`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
		err := kafka.SendMsg(kafkaURL, topic, numMsg, targetOffset, batchSize)
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
	sendMsgCmd.PersistentFlags().Int64Var(&numMsg, "num-msg", 0, "Number of messages to send. Cannot set together with target-offset")
	sendMsgCmd.PersistentFlags().Int64Var(&targetOffset, "target-offset", 0, "Send message until the topic reaches the target offset. Cannot set together with num-msg")
	sendMsgCmd.PersistentFlags().Int64Var(&batchSize, "batch-size", 1000, "Size of batch when sending.")
}
