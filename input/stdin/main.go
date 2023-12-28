package stdin

import (
	"fmt"
	"strconv"
)

type StdInput struct {
}

func (s *StdInput) ReadOneInput() (int, bool,error) {
	var userInput string

	fmt.Print("Enter something: ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return 0, true, nil
	}

	intNumber, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, false, err
	}
	return intNumber, false, nil
}
