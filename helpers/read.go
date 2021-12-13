package helpers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type FileToRead struct {
	Filename    string
	CustomLogic func(readValue string) error
}
type ReadValues interface {
	ReadValues() error
}

func (r FileToRead) ReadValues() error {
	xb, err := ioutil.ReadFile(r.Filename)
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
