package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

func DoWork(bufferSize int, producer *Producer, id int) {
	buffer := make([]byte, bufferSize)
	sha := sha256.New()
	for {
		isMore, bytesRead, index := producer.GetMoreJob(buffer)
		if !isMore {
			break
		}

		computedHash := computeHash(sha, buffer[:bytesRead])
		fmt.Printf("worker %v batch %d: %x\n", id, index, computedHash)
	}
}

func computeHash(sha hash.Hash, buffer []byte) []byte {
	return sha.Sum(buffer)
}
