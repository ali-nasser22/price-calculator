package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		err := file.Close()
		if err != nil {
			return nil, errors.New("error closing file")
		}
		return nil, errors.New("error scanning file")
	}
	err = file.Close()
	if err != nil {
		return nil, errors.New("error closing file")
	}
	return lines, nil
}
