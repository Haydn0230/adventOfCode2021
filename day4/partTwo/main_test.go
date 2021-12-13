package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = `79,88,67,56,53,97,46,29,37,51,3,93,92,78,41,22

17 50 13 93 20
68 57 76 92 86
 2 91 67 78 11
94 70 84 41 25
32 90 45 22 41

78 27 82 68 20
14  2 34 51  7
58 57 99 37 81
 9  4  0 76 45
67 69 70 17 23

38 60 62 34 41
39 58 91 45 10
66 74 94 50 17
68 27 75 97 49
36 64  5 98 15
`

type TestFileToRead struct {
	Filename    string
	CustomLogic func(readValue string) error
}

func Test_readTables(t *testing.T) {
	emptyMap := make(map[int]table)
	d := splitCSVAndWhitespaceData{
		tables: emptyMap,
	}

	testFileToRead := TestFileToRead{
		Filename: "test",
		CustomLogic: d.splitCSVAndWhitespaceInput,
	}

	err := readTables(testFileToRead)
	if err != nil {
		t.Error(err)
	}

	answer := lastWinner(d.callList, d.tables )

	fmt.Printf("The answer is %v:", answer)
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

func Test_removeKey(t *testing.T) {
	keyStore = []int{1,2,3,4,5,6}
	expected := []int{1,2,3,4,5}
	removeKey(6)
	assert.Equal(t, expected, keyStore)

}