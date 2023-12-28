package input

import (
	"golang-interface/input/file"
	"golang-interface/input/stdin"
	"os"
)

type Input interface {
	ReadOneInput() (int, bool, error)
}

func Init() (Input, error) {
	// Access the environment variables
	stdCase := os.Getenv("STDIN_CASE")
	var input Input
	if stdCase == "Y" {
		tmp := stdin.StdInput{}
		input = &tmp
	} else {
		tmp, err := file.Init()
		if err != nil {
			return nil, err
		}
		input = &tmp
	}
	return input, nil
}
