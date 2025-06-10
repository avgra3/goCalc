package calc

import (
	"errors"
	"math"
)

func PeriodOverPeriod(current, previous int64) (float64, error) {
	if previous == 0 {
		return 0, errors.New("Previous period cannot be zero")
	}
	if previous < 0 {
		return 0, errors.New("Previous period cannot have a value less than or equal to zero")
	}
	if current < 0 {
		return 0, errors.New("Current period cannot have a value less than or equal to zero")
	}

	value, err := percentChange(float64(previous), float64(current))
	// Rounds our result to 2 decimals
	return roundFloats(value, 2), err
}

func percentDifference(value1, value2 float64) (float64, error) {
	if value1 < 0 || value2 < 0 {
		return 0, errors.New("value1 and value2 must be greater than 0")
	}

	result := 100 * (math.Abs(value1-value2) / ((value1 + value2) / 2))
	return result, nil
}

func percentChange(value1, value2 float64) (float64, error) {
	if value1 < 0 || value2 < 0 {
		return 0, errors.New("value1 and value2 must be greater than 0")
	}

	result := 100 * ((value2 - value1) / math.Abs(value1))
	return result, nil

}

func roundFloats(value float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio
}
