package fileops // To be usable like this, the fileops.go needs to be put inside a folder named 'fileops'

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// In Go, function whose name is capitalized are the ones that gets exported. So GetFloatFromFile is accessible by other files using the same package, while getFloatFromFile is not.

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

