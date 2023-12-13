package main

import (
	"io"
	"os"
	"sync"
)

type Producer struct {
	file  *os.File
	index int
	mutex sync.Mutex
}

func (producer *Producer) GetMoreJob(buffer []byte) (isRead bool, bytesRead int, batchIndex int) {
	producer.mutex.Lock()
	defer producer.mutex.Unlock()

	bytesRead, err := producer.file.Read(buffer)
	if err != nil {
		if err == io.EOF {
			return false, 0, 0
		} else {
			panic(err)
		}
	}
	producer.index++
	return true, bytesRead, producer.index
}
