package cmd

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	kafkaURL, cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafka-client",
	Short: "Interact with kafka quickly",
	Long:  `kafka-client helps you to interact with kafka efficiently, including list topic, get offset or send messages`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	start := time.Now() // record the time of writing messages
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	timeElapsed := time.Since(start).String()
	log.Infof("Execution time is %s\n", timeElapsed)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kafka-client.yaml)")
	rootCmd.PersistentFlags().StringVar(&kafkaURL, "kafka-url", "localhost:9092", "The kafka broker URL to connect with.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "kafka-client is an efficient client to interact with kafka cluster")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kafka-client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kafka-client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
