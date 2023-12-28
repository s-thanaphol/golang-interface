package filelog

import (
	"io/ioutil"
	"strconv"
)

type FileLog struct {
	number int
}

func byteArrayToInt(byteArray []byte) (int, error) {
	strNumber := string(byteArray)
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func readFile() (int, error) {
	content, err := ioutil.ReadFile("./database/filelog/number.txt")
	if err != nil {
		return 0, err
	}
	return byteArrayToInt(content)
}

func writeFile(content string) error {
	err := ioutil.WriteFile("./database/filelog/number.txt", []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileLog) Get() (int, error) {
	number, err := readFile()
	if err != nil {
		return 0, err
	}
	f.number = number
	return number, nil
}

func (f *FileLog) Post(number int) error {
	err := writeFile(strconv.Itoa(number))
	if err != nil {
		return err
	}
	f.number = number
	return nil
}