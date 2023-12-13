package args

import (
	"fmt"
	"os"
	"strconv"
)

type Arguments struct {
	FilePath  string
	BlockSize int
}

func ParseCmdArguments() Arguments {
	cmd := os.Args[:1]

	if len(cmd) != 2 {
		return getDefault()
	}

	blockSize, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Printf("%v is not a parsable number\n", cmd[1])
		return getDefault()
	}

	args := Arguments{cmd[0], blockSize}
	return args
}

func getDefault() Arguments {
	return Arguments{"default.txt", 40}
}
