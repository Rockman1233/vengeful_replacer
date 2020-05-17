package scanners

import (
	"bufio"
	"errors"
	"os"
)

type StdinScanner struct {
	file    os.File
	message string
}

func (scn StdinScanner) GetData() string {
	return scn.message
}

func (scn *StdinScanner) SetSource(file os.File) error {

	scn.file = file

	_, err := file.Stat()

	// If empty data or error of reading
	if err != nil {
		return errors.New("empty source")
	}

	bufioScanner := bufio.NewScanner(&file)
	bufioScanner.Scan()

	scn.message = bufioScanner.Text()

	return nil
}

func NewStdinScanner() *StdinScanner {

	newEntity := StdinScanner{}

	return &newEntity
}
