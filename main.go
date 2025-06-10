package main

import (
	"fmt"
	"log"
	"math"

	"github.com/avgra3/goCalc/internal/calc"
)

func main() {
	var current int64
	var previousPeriod int64
	var previousYear int64

	fmt.Print("Current period count: ")
	_, err := fmt.Scanln(&current)
	if err != nil {
		log.Fatalf("ERROR! %v", err)
	}
	fmt.Print("Previous period count: ")
	_, err = fmt.Scanln(&previousPeriod)
	if err != nil {
		log.Fatalf("ERROR! %v", err)
	}
	fmt.Print("Previous year count: ")
	_, err = fmt.Scanln(&previousYear)
	if err != nil {
		log.Fatalf("ERROR! %v", err)
	}

	pop, err := calc.PeriodOverPeriod(current, previousPeriod)
	if err != nil {
		log.Fatalf("There was an error! %v\n", err)
	}
	yoy, err := calc.PeriodOverPeriod(current, previousYear)
	if err != nil {
		log.Fatalf("There was an error! %v\n", err)
	}

	messagePOP := fmt.Sprintf("PoP: %v", message(pop))
	messageYOY := fmt.Sprintf("YoY: %v", message(yoy))

	fmt.Println(messagePOP)
	fmt.Println(messageYOY)
}

func message(value float64) string {
	if value == 0 {
		return fmt.Sprintf("No change (0%%)")
	}
	if value < 0 {
		if value > -1 {
			return fmt.Sprintf("Decrease <1%%")
		}
		return fmt.Sprintf("Decrease of %v%%", math.Abs(value))
	}
	if value > 0 && value < 1 {
		return fmt.Sprintf("Increase <1%%")
	}
	return fmt.Sprintf("Increase of %v%%", value)
}
