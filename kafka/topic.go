package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// ListTopics will list existing topics at the broker
func ListTopics(kafkaURL string) error {
	var err error
	var conn *kafka.Conn
	var names []string

	conn, err = kafka.Dial("tcp", kafkaURL)
	if err != nil {
		return err
	}
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

	conn, err = kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, 0)
	if err != nil {
		return err
	}
	defer conn.Close()

	lastOffsetBefore, err := getLastOffset(conn)
	if err != nil {
		return err
	}
	fmt.Printf("Topic %s's offset = %d\n", topic, lastOffsetBefore)

	return nil
}
