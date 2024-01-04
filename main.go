package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func convertTimeFormat(inputTime string) (string, error) {
	tt, err := time.Parse("03:04:05PM", inputTime)
	if err != nil {
		return "", fmt.Errorf("%w: %w", "unable to parse error", err)
	}

	return tt.Format(time.TimeOnly), nil
}

func main() {
	flag.Parse()
	timeInput := flag.Args()

	if len(timeInput) != 1 {
		fmt.Printf("Expecting one value, got %d instead\n", len(timeInput))
		os.Exit(1)
	}

	//converting args I got into 24H format
	convertedTime, err := convertTimeFormat(timeInput[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(convertedTime)

}
