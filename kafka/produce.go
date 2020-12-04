package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// SendMsg will send messages to the kafka in batches
func SendMsg(kafkaURL, topic string, numMsg, targetOffset, batchSize int64) error {
	var err error
	var connLeader *kafka.Conn
	var numMsgSend int64

	log.Infof("SendMsg with params: kafkaURL=%s, topic=%s, numMsg=%d, batchSize=%d, targetOffset=%d", kafkaURL, topic, numMsg, batchSize, targetOffset)
	if kafkaURL == "" || topic == "" || (numMsg <= 0 && targetOffset <= 0) || (numMsg > 0 && targetOffset > 0) || batchSize <= 0 {
		return errors.New("input parameters are not valid")
	}

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
	log.Infof("Before sending, the offset = %d\n", lastOffsetBefore)

	if numMsg > 0 { // will ignore targetOffset
		numMsgSend = numMsg
	} else { // Should use the targetOffset
		numMsgSend = targetOffset - lastOffsetBefore
	}

	if numMsgSend <= 0 { // targetOffset <= lastOffsetBefore
		log.Infof("No need to send msg with targetOffset=%d, currentOffset=%d\n", targetOffset, lastOffsetBefore)
		return nil
	}

	batchMsgs := newBatchMsg(batchSize)                               // construct a batch message
	remainMsgs := newBatchMsg(numMsgSend % batchSize)                 // construct a batch message for remaining messages
	_ = connLeader.SetWriteDeadline(time.Now().Add(30 * time.Second)) // set sending timeout

	// send batch messages
	for i := int64(0); i < numMsgSend/batchSize; i++ {
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
	log.Infof("After sending %d messages, the offset = %d\n", numMsgSend, lastOffsetAfter)

	if lastOffsetAfter-lastOffsetBefore != numMsgSend {
		return fmt.Errorf("targeted numMsg=%d, but kafka received %d msgs", numMsg, lastOffsetAfter-lastOffsetBefore)
	}
	return nil
}
