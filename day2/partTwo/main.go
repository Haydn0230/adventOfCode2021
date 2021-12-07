package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main(){
	pathHistory := make([]string, 0)
	err := ReadValues("input.txt", func(readValue string) error {
		pathHistory = append(pathHistory, readValue)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

var (
	horizontal = 0
	depth = 0
	aim = 0
)

	for _, path := range pathHistory {
		splitPath := strings.Split(path, " ")
		distance, err := strconv.Atoi(splitPath[1])
		if err !=nil {
			fmt.Println(err)
		}

		switch splitPath[0] {
		case "forward":
			horizontal += distance
			depth += distance * aim
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}


	fmt.Println(horizontal * depth)
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