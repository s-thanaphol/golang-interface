package file

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type FileInput struct {
	data []int
}

func (f *FileInput) ReadOneInput() (int, bool, error) {
	if len(f.data) == 0 {
		return 0, true, nil
	}
	res := f.data[0]
	f.data = f.data[1:]

	return res, false, nil
}

func Init() (FileInput, error) {
	filePath := "./input.txt"
	res := FileInput{data: []int{}}
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return res, err
	}

	fileContent := string(content)
	fmt.Println(fileContent)

	// Split the content by newline
	lines := strings.Split(fileContent, "\n")
	fmt.Println(lines)

	// Print each line
	data := make([]int, len(lines))
	for i, line := range lines {
		fmt.Println(i, line)
		intNumber, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return res, err
		}
		data[i] = intNumber
	}
	res.data = data

	return res, nil
}
