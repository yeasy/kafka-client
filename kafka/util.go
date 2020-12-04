package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

// newBatchMsg create a message array with given batchSize
func newBatchMsg(batchSize int64) []kafka.Message {

	var batchMsgs []kafka.Message
	for i := int64(0); i < batchSize; i++ {
		batchMsgs = append(batchMsgs, kafka.Message{Value: []byte("")})
	}

	return batchMsgs
}

// getLastOffset fetches the latest offset from kakfa
func getLastOffset(connLeader *kafka.Conn) (int64, error) {
	lastOffset, err := connLeader.ReadLastOffset()
	if err != nil {
		return 0, fmt.Errorf("failed to get the latest offset with err=%v", err)
	}
	return lastOffset, nil
}
