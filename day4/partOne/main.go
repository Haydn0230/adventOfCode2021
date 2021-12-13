package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Haydn0230/advent2021/helpers"
)


type table struct {
	column [][]cell
	row [][]cell
}

type cell struct {
	value int
	marked bool
}

type splitCSVAndWhitespaceData struct {
	callList     []int
	tables       map[int]table
	tableCounter int
	rowCounter   int
}

// this is over engineered to use interfaces. Classic example of we can but should we?
func main(){
	emptyMap := make(map[int]table)
	d := splitCSVAndWhitespaceData{
		tables: emptyMap,
	}

	file := helpers.FileToRead{
		Filename:"input.txt",
		CustomLogic: d.splitCSVAndWhitespaceInput,

	}
	// read the tables
	err := readTables(file)
	if err != nil {
		panic(err)
	}

	answer := playBingo(d.tables, d.callList)

	fmt.Printf("The answer is %v:", answer)
}
func winner(numberCalled int, valuesToCheck []cell) ([]cell, bool) {
		markedCounter := 0
		for i, cell  := range valuesToCheck {
			if cell.value == numberCalled {
				cell.marked = true
				valuesToCheck[i] = cell
				markedCounter ++
				continue
			}

			if cell.marked {
				markedCounter ++
			}

		}
		if markedCounter == len(valuesToCheck) {
			return valuesToCheck, true
		}
		return valuesToCheck, false

}

func sumUnmarked(valuesToCheck [][]cell) int {
	total := 0
	for _, row := range valuesToCheck {
		for _, cell := range row {
			if !cell.marked {
				total += cell.value
			}
		}
	}

	return total
}


func playBingo(tables map[int]table, callList []int) int {
	for _, numberCalled := range callList {
		for _, table := range tables {
			// check rows
			for _, row := range table.row {
				r, isWinner := winner(numberCalled, row)
				// update row with the checked values
				row = r
				if isWinner {
					return numberCalled * sumUnmarked(table.row)
				}
			}

			// check columns
			for _, column := range table.column {
				c, isWinner := winner(numberCalled, column)
				// update row with the checked values
				column = c
				if isWinner {
					return numberCalled * sumUnmarked(table.row)
				}
			}
		}
	}
	return 0
}



func readTables(file helpers.ReadValues) error {
	// read the tables
	err := file.ReadValues()
	if err != nil {
		return err
	}
	return nil
}

func (d *splitCSVAndWhitespaceData) splitCSVAndWhitespaceInput(readValue string) error {
	// read the callList to be shouted out
	csv := strings.Split(readValue, ",")
	if len(csv) > 1 {
		var err error
		d.callList, err = readCSV(csv)
		if err != nil {
			panic(err)
		}
		return nil
	}

	// we've hit a line break
	if len(readValue) == 0 {
		// increment table counter as we're onto a new table
		d.tableCounter ++
		return nil
	}

	// read the tables

	// replace double spaces with single to ensure we don't have blank strings when we split
	singleSpacedValue := strings.ReplaceAll(readValue, "  ", " ")
	rawData := strings.Split(singleSpacedValue, " ")

	tempRow := make([]cell, 0, len(rawData))
	tempTable := d.tables[d.tableCounter]
	for columnIdx, valueAsString := range rawData {
		valueAsInt, err := strconv.Atoi(valueAsString)
		if err != nil {
			return err
		}
		// create new cell with value
		newCell := cell{
			value:  valueAsInt,
			marked: false,
		}

		// add the cell to the row
		tempRow = append(tempRow, newCell)

		// check if the column exists first
		if tempTable.column == nil || len(tempTable.column) == columnIdx {
			tempTable.column = append(tempTable.column,[]cell{newCell})
		} else {
			tempTable.column[columnIdx] = append(tempTable.column[columnIdx], newCell)
		}

	}

	tempTable.row = append(tempTable.row, tempRow)
	d.rowCounter ++
	d.tables[d.tableCounter] = tempTable
	return nil
}

func readCSV(csv []string) ([]int, error) {
 n := make([]int, 0)
	for _, valueAsString := range csv {
		if valueAsString != "" {
			valueAsInt, err := strconv.Atoi(valueAsString)
			if err != nil {
				return []int{}, err
			}
			n = append(n, valueAsInt)
		}
	}
	return n, nil
}