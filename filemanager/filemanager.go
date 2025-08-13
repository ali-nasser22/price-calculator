package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath string, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
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

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
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
