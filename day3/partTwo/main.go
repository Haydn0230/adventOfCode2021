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
	data := make([]string, 0)
	//data := []string{
	//	"000",
	//	"000",
	//	"100",
	//	"100",
	//	"100", // <- lowest
	//
	//}
	err := ReadValues("input.txt", func(readValue string) error {
		data = append(data, readValue)
		return nil
	})
	if err != nil {
		fmt.Println("err", err)
	}

	dataSplitByPosition, err := splitByPosition(data)
	if err != nil {
		fmt.Println(err)
	}

	lifeSupportRating, err := lifeSupport(data)
	if err != nil {
		fmt.Println(err)
	}
	energy, err := calcEnergy(dataSplitByPosition)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Printf("Energy %v\n", energy)
	fmt.Printf("Life Support %v\n", lifeSupportRating)
}

func splitByPosition(binaryInput []string) (map[int][]int, error) {
	binaryOutput := make(map[int][]int)
	for _, line := range binaryInput {
		binaryDigits := strings.Split(line, "")
		for i, b := range binaryDigits {
			convertedDigit, err := strconv.Atoi(b)
			if err != nil {
				return nil, err
			}
			binaryOutput[i] = append(binaryOutput[i], convertedDigit)
		}
	}

	return binaryOutput, nil

}

func lifeSupport(data []string) (int, error) {
	C02, err := findNumberByFrequency(false, data)
	if err != nil {
		return 0, err
	}
	oxygen, err := findNumberByFrequency(true, data)
	if err != nil {
		return 0, err
	}

	return int(oxygen * C02), nil
}

func findNumberByFrequency(highlyFrequent bool, data []string) (int64, error) {
	currentList := data
	tempList := make([]string, 0)
	keys := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for _, k := range keys {
		tempList = []string{}
		if len(currentList) == 1 {
			fmt.Println("got it", currentList)
			return strconv.ParseInt(currentList[0], 2, 32)
		}
		dataSplitByPosition, err := splitByPosition(currentList)
		if err != nil {
			return 0, err
		}
		// decide if we want to calculate oxygen or C02
		highFrequency, lowFrequency := calcHighLow(dataSplitByPosition[k], highlyFrequent)
		frequency := lowFrequency
		if highlyFrequent {
			frequency = highFrequency
		}

		for _, byteLine := range currentList {
			byts := strings.Split(byteLine, "")
			for i, b := range byts {
				if i == k {
					convertedDigit, err := strconv.Atoi(b)
					if err != nil {
						fmt.Println(err)
					}
					if convertedDigit == frequency {
						tempList = append(tempList, byteLine)
					}
				}

			}
		}
		currentList = tempList
	}

	if len(currentList) == 1 {
		return strconv.ParseInt(currentList[0], 2, 32)
	}

	return 0, fmt.Errorf("could not find value")
}

func calcEnergy(dataOutput map[int][]int) (int, error) {
	var (
		gamma   = ""
		epilson = ""
	)

	keys := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, k := range keys {
		highFrequency, lowFrequency := calcHighLow(dataOutput[k], true)
		gamma = fmt.Sprintf("%v%v", gamma, strconv.Itoa(highFrequency))
		epilson = fmt.Sprintf("%v%v", epilson, strconv.Itoa(lowFrequency))
		fmt.Println(highFrequency, lowFrequency)
	}

	e, err := strconv.ParseInt(epilson, 2, 32)
	if err != nil {
		return 0, err
	}

	g, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		return 0, err
	}
	fmt.Printf("epilson: %v\n  gamma: %v\n", e, g)
	fmt.Printf("epilson: %v\n  gamma: %v\n", epilson, gamma)
	return int(e * g), nil
}

func calcHighLow(binaryInput []int, highlyFrequent bool) (int, int) {
	//fmt.Println(key, binaryInput)
	zero := make([]int, 0)
	one := make([]int, 0)

	for _, v := range binaryInput {
		if v == 0 {
			zero = append(zero, v)
		} else if v == 1 {
			one = append(one, v)
		}
	}
	if len(zero) > len(one) {
		//fmt.Println("zero is highest")
		return 0, 1
	}

	//if len(zero) == len(one) {
	//	if !highlyFrequent {
	//		return 0, 1
	//	}
	//}

	return 1, 0
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
