package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {

	floats := make([]float64, len(strings))

	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil

}
