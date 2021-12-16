package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

var testData = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

type testCoords struct {

}

func Test(t *testing.T) {
	d := coordsData{
		ventCords:  make(map[int]coords),
		mapCounter: 0,
	}

 	f := TestFileToRead{
		Filename:    "",
		CustomLogic: d.coordinates,
	}

	err := f.ReadValues()
	if err != nil {
		panic(err)
	}

	vm := createVentMapHorizontalAndVertical(d.ventCords)

	for _, v := range vm.ventLines {
		fmt.Println(v)
	}

	fmt.Println(crossPoints(vm))

}

type TestFileToRead struct {
	Filename    string
	CustomLogic func(readValue string) error
}

func (r TestFileToRead) ReadValues() error {
	buf := bytes.NewBufferString(testData)

	for {
		line, readError := buf.ReadString('\n')
		if readError != io.EOF {

		}
		line = strings.TrimSpace(
			strings.Replace(line, "\n", " ", 1),
		)

		err := r.CustomLogic(line)
		if err != nil {
			return err
		}

		if readError == io.EOF {
			break
		}
	}

	return nil
}