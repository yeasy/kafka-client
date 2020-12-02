package kafka

import (
	"fmt"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// Deprecated
func getParamsFromEnv() (string, string, int, int) {
	kafkaURL := os.Getenv("kafkaURL")
	if kafkaURL == "" {
		kafkaURL = "localhost:9092"
	}

	fmt.Printf("kafkaURL=%s\n", kafkaURL)
	topic := os.Getenv("topic")
	if topic == "" {
		topic = "test"
	}
	fmt.Printf("topic=%s\n", topic)
	numMsgStr := os.Getenv("numMsg")
	if numMsgStr == "" {
		numMsgStr = "10000"
	}
	numMsg, err := strconv.Atoi(numMsgStr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("numMsg=%d\n", numMsg)

	batchSizeStr := os.Getenv("batchSize")
	if batchSizeStr == "" {
		batchSizeStr = "1000"
	}
	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("batchSize=%d\n", batchSize)

	return kafkaURL, topic, numMsg, batchSize
}

// newBatchMsg create a message array with given batchSize
func newBatchMsg(batchSize int) []kafka.Message {

	var batchMsgs []kafka.Message
	for i := 0; i < batchSize; i++ {
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
