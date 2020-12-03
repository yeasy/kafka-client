package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// SendMsg will send messages to the kafka in batches
func SendMsg(kafkaURL, topic string, numMsg, batchSize int) error {
	var err error
	var connLeader *kafka.Conn

	connLeader, err = kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, 0)
	if err != nil {
		return err
	}
	log.Debugf("Connected to kafka %s", kafkaURL)
	defer connLeader.Close()

	lastOffsetBefore, err := getLastOffset(connLeader)
	if err != nil {
		return err
	}
	log.Infof("Before sending messages, the last offset = %d\n", lastOffsetBefore)

	batchMsgs := newBatchMsg(batchSize)                               // construct a batch message
	remainMsgs := newBatchMsg(numMsg % batchSize)                     // construct a batch message for remaining messages
	_ = connLeader.SetWriteDeadline(time.Now().Add(30 * time.Second)) // set sending timeout

	// send batch messages
	for i := 0; i < numMsg/batchSize; i++ {
		_, err = connLeader.WriteMessages(batchMsgs...)
		if err != nil {
			log.Errorf("failed to write batch messages: %d with error=%v", i, err)
		}
		if i*batchSize%100000 == 0 {
			log.Debugf("Sent %d batched messages to kafka, with batchSize=%d", i, batchSize)
		}
	}
	// send remaining messages in a batch
	_, err = connLeader.WriteMessages(remainMsgs...)
	if err != nil {
		log.Errorf("failed to write remaining messages with error=%v", err)
	}

	lastOffsetAfter, err := getLastOffset(connLeader)
	if err != nil {
		return err
	}
	log.Infof("After sending %d messages, the last offset = %d\n", numMsg, lastOffsetAfter)

	if int(lastOffsetAfter-lastOffsetBefore) != numMsg {
		return fmt.Errorf("targeted numMsg=%d, but kafka received %d msgs", numMsg, lastOffsetAfter-lastOffsetBefore)
	}
	return nil
}
