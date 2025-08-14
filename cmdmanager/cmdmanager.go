package cmdmanager

import (
	"errors"
	"fmt"
)

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your Prices. Confirm price with [ENTER]")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		_, err := fmt.Scan(&price)
		if err != nil {
			return nil, errors.New("error reading your Price")
		}
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
