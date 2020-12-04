package kafka

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/segmentio/kafka-go"
)

// ListTopics will list existing topics at the broker
func ListTopics(kafkaURL string) error {
	var err error
	var conn *kafka.Conn
	var names []string

	log.Infof("GetOffset with params: kafkaURL=%s", kafkaURL)
	conn, err = kafka.Dial("tcp", kafkaURL)
	if err != nil {
		return err
	}
	log.Debugf("Connected to kafka %s", kafkaURL)
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	t := map[string]struct{}{}

	for _, p := range partitions {
		if _, ok := t[p.Topic]; !ok {
			names = append(names, p.Topic)
			t[p.Topic] = struct{}{}
		}
	}
	fmt.Printf("%d topics: %v\n", len(names), names)

	return nil
}

// GetOffset will get the last offset in the topic
func GetOffset(kafkaURL, topic string) error {
	var err error
	var conn *kafka.Conn

	log.Infof("GetOffset with params: kafkaURL=%s, topic=%s", kafkaURL, topic)
	conn, err = kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, 0)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Debugf("Connected to kafka %s", kafkaURL)

	lastOffsetBefore, err := getLastOffset(conn)
	if err != nil {
		return err
	}
	fmt.Printf("Topic %s's offset = %d\n", topic, lastOffsetBefore)

	return nil
}
