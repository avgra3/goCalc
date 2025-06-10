package calc

import (
	"math"
	"testing"
)

func TestPeriodOverPeriod(t *testing.T) {
	currents := []int64{}
	previous := []int64{}

	// Should have no errors
	for i := range currents {
		_, err := PeriodOverPeriod(currents[i], previous[i])
		if err != nil {
			t.Error(err)
		}
	}

	// Ensure errors
	currentErrors := []int64{10, 10, -10}
	previousErrors := []int64{0, -10, 10}
	// previous == 0
	_, err := PeriodOverPeriod(currentErrors[0], previousErrors[0])
	if err == nil {
		t.Errorf("Expect a value of 0 for previous to error")
	}
	// previous < 0
	_, err = PeriodOverPeriod(currentErrors[1], previousErrors[1])
	if err == nil {
		t.Errorf("Expect a value less than 0 for previous to error")
	}
	// current < 0
	_, err = PeriodOverPeriod(currentErrors[2], previousErrors[2])
	if err == nil {
		t.Errorf("Expect a value less than 0 for current to error")
	}

}

func TestPercentDifference(t *testing.T) {
	// Tests accuracy
	inputs := [][]float64{
		{10, 10},
		{123, 456},
		{192, 134},
		{121212, 121212212},
	}
	expected := []float64{0.0, 115.026, 35.5828, 199.6}
	for i := range inputs {
		percentDiff, _ := percentDifference(inputs[i][0], inputs[i][1])
		if !floatEqualRelative(expected[i], percentDiff) {
			t.Errorf("Expected inputs v1=%f and v2=%f, result to equal expected (they did not). %f (calculated) != %f (expected)", inputs[i][0], inputs[i][1], percentDiff, expected[i])
		}
	}
}

func TestPercentChange(t *testing.T) {
	// Tests accuracy
	inputs := [][]float64{
		{10, 10},
		{123, 456},
		{192, 134},
	}
	expected := []float64{0.0, 270.732, -30.2083}
	for i := range inputs {
		percentDiff, _ := percentChange(inputs[i][0], inputs[i][1])
		if !floatEqualRelative(expected[i], percentDiff) {
			t.Errorf("Expected inputs v1=%f and v2=%f, result to equal expected (they did not). %f (calculated) != %f (expected)", inputs[i][0], inputs[i][1], percentDiff, expected[i])
		}
	}
}

func TestRoundFloats(t *testing.T) {
	values := []float64{0.0, 100.000001, 1000.000001}
	precision := []uint{1, 2, 3}
	expected := []float64{0.0, 100.0, 1000.00}

	for i := range values {
		rounded := roundFloats(values[i], precision[i])
		if rounded != expected[i] {
			t.Errorf("Expected rounded (%f) == expected (%f) (it does not)", rounded, expected[i])
		}
	}
}

// A helper to compare floats
func floatEqualRelative(a, b float64) bool {
	tolerance := 0.00001
	if a == b {
		return true
	}
	diff := math.Abs(a - b)
	mean := math.Abs(a+b) / 2.0
	return (diff / mean) < tolerance
}
