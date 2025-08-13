package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJson(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed creating file")
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		err := file.Close()
		if err != nil {
			return errors.New("failed closing file")
		}
		return err
	}
	err = file.Close()
	if err != nil {
		return errors.New("failed closing file")
	}
	return nil
}
