package cmd

import (
	"kafka-client/kafka"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getOffsetCmd represents the getOffset command
var getOffsetCmd = &cobra.Command{
	Use:   "getOffset",
	Short: "Get the last offset in the topic",
	Long:  `Get the last offset in the topic`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
		err := kafka.GetOffset(kafkaURL, topic)
		if err != nil {
			log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(getOffsetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getOffsetCmd.PersistentFlags().String("foo", "", "A help for foo")
	getOffsetCmd.PersistentFlags().StringVar(&topic, "topic", "test", "The kafka topic to check the offset.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getOffsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
