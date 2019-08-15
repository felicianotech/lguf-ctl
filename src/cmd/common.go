package cmd

import (
	"fmt"
	"strconv"

	"github.com/felicianotech/go-lguf/lguf"
)

// coverts a percentage string or integer brightness level to uint16
func prepBrightnessInput(rawValue string) (uint16, error) {

	var value uint16
	pctIndex := len(rawValue) - 1

	// check for a percentage
	if rawValue[pctIndex:] == "%" {

		intValue, err := strconv.Atoi(rawValue[0:pctIndex])
		if err != nil {
			return 0, fmt.Errorf("Percentage contains invalid characters.")
		}

		if intValue < 0 || intValue > 100 {
			return 0, fmt.Errorf("Percentage should be between 0-100%.")
		}

		// convert percentage to integer values
		value = uint16(float64(lguf.MaxBrightness) * (float64(intValue) / 100.00))
	} else {

		intValue, err := strconv.Atoi(rawValue)
		if err != nil {
			return 0, fmt.Errorf("Value was not an integer.")
		}

		value = uint16(intValue)
	}

	return value, nil
}
