package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	container := make([]int, 0)
	returnInts := func(stringToConvert string) error {
		i, err := strconv.Atoi(stringToConvert)
		if err != nil {
			return fmt.Errorf("failed to convert to int: %v", err)
		}

		container = append(container, i)

		return nil
	}

	err := ReadValues("input.txt", returnInts)
	if err != nil {
		fmt.Printf("error reading values %v", err)
	}

	var previousV int
	var results []string
	for _, currentV := range container {
		if previousV != 0 && previousV < currentV {
			results = append(results, "Increased")
		}
		previousV = currentV
	}

	fmt.Printf("number increased - %v\n", len(results))
}

func ReadValues(filename string, customLogic func(readValue string) error) error {
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
