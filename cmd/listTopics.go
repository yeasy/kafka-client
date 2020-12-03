package cmd

import (
	"kafka-client/kafka"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// listTopicsCmd represents the listTopics command
var listTopicsCmd = &cobra.Command{
	Use:   "listTopics",
	Short: "List the existing topics at the broker",
	Long:  `List the existing topics at the broker`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
		err := kafka.ListTopics(kafkaURL)
		if err != nil {
			log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(listTopicsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTopicsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTopicsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
