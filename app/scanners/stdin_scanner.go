package scanners

import (
	"bufio"
	"os"
)

type StdinScanner struct {
	file    os.File
	message string
}

func (scn StdinScanner) GetData() string {
	return scn.message
}

func NewStdinScanner(file os.File) (StdinScanner, error) {

	newEntity := StdinScanner{file: file}

	_, err := file.Stat()

	// If empty data or error of reading
	if err != nil {
		return newEntity, err
	}

	bufioScanner := bufio.NewScanner(&file)
	bufioScanner.Scan()
	newEntity.message = bufioScanner.Text()

	return newEntity, nil
}
