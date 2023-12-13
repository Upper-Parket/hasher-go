package main

import (
	"FileHasher/args"
	"fmt"
	"os"
	"sync"
)

func main() {
	arguments := args.ParseCmdArguments()

	file, err := os.Open(arguments.FilePath)
	if err != nil {
		fmt.Printf("error opening file %v: %v\n", arguments.FilePath, err)
		return
	}

	fmt.Printf("opened file named %v\n", file.Name())
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error occurred while closing file %v\n", err)
		}
	}(file)

	producer := Producer{file: file}
	var waitGroup sync.WaitGroup
	for id := 1; id <= 4; id++ {
		waitGroup.Add(1)

		go func(id int) {
			defer waitGroup.Done()
			DoWork(arguments.BlockSize, &producer, id)
		}(id)
	}
	waitGroup.Wait()
	fmt.Println("\ndone")
}
