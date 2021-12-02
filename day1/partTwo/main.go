package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var rollingAverage = 3

var input = "input.txt"

func main() {
	container := make([]int, 0)
	returnInts := func(stringToConvert string) error{
		i, err := strconv.Atoi(stringToConvert)
		if err !=nil {
			return fmt.Errorf("failed to convert to int: %v", err)
		}

		container = append(container, i)

		return nil
	}

	err := ReadValues(input, returnInts)
	if err != nil {
		fmt.Printf("error reading values %v", err)
	}

	window :=make([]int, 0, rollingAverage)
	var results []int
	for _, currentV  := range container {
		window = append(window,currentV)
		if len(window) <=rollingAverage {
			continue
		}

		previousW := sumIntArray(window[:rollingAverage])
		currentW := sumIntArray(window[1:])
		if previousW != 0 && previousW < currentW {
			results = append(results, 1)
		}

		window = window[1:]
	}

	fmt.Printf("number increased - %v\n", len(results))
}

func sumIntArray(values []int) int {
	var total int
	for _, v := range values {
		total += v
	}
	return total
}


func ReadValues(filename string, customLogic func(readValue string)error)  error {
	xb, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file - %v", err)
	}

	buf := bytes.NewBuffer(xb)

	for {
		line, readError := buf.ReadString('\n')
		if readError != io.EOF {

		}
		line = strings.TrimSpace(
			strings.Replace(line, "\n", " ", 1),
			)

		err := customLogic(line)
		if err != nil {
			return err
		}

		if readError == io.EOF {
			break
		}
	}
	return nil
}