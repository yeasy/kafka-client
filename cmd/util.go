package cmd

import (
	log "github.com/sirupsen/logrus"
)

func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Warnf("Wrong log level specified with %s, will use default INFO level", logLevel)
		log.SetLevel(log.InfoLevel)
	}
}
