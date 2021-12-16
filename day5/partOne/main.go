package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Haydn0230/advent2021/helpers"
)


type coords struct {
	x []int
	y []int
}

type coordsData struct {
	ventCords  map[int]coords
	mapCounter int
}

// [][]int first [] is y value second [] is x
type ventMap struct {
	ventLines [][]string
}

func main() {
	d := coordsData{
		ventCords: make(map[int]coords),
	}
	f := helpers.FileToRead{
		Filename:    "input.txt",
		CustomLogic: d.coordinates,
	}
	err := f.ReadValues()
	if err != nil {
		panic(err)
	}

	mapOfVents := createVentMapHorizontalAndVertical(d.ventCords)
	crossOver := crossPoints(mapOfVents)

	fmt.Printf("the number of cross over points %v\n", crossOver)
}

func crossPoints(ventMapData ventMap) int {
	var crossPointCounter int
	for _, vertical := range ventMapData.ventLines {
		for _, horizontal := range vertical {
			if horizontal == "." {
				continue
			}
			i, err := strconv.Atoi(horizontal)
			if err != nil {
				fmt.Println(err)
			}

			if i > 1 {
				crossPointCounter ++
			}
		}
	}

	return crossPointCounter
}

func createVentMapHorizontalAndVertical(ventCoords map[int]coords) ventMap {
	keys := make([]int, 0, len(ventCoords))
	for k :=  range ventCoords {
		keys = append(keys, k)
	}

	ventLocation := ventMap{
		ventLines: make([][]string, 0),
	}
	sort.Ints(keys)

	for _, k := range keys {
		sort.Ints(ventCoords[k].x)
		sort.Ints(ventCoords[k].y)
		for y:= ventCoords[k].y[0]; y <= ventCoords[k].y[1]; y ++ {
			for  len(ventLocation.ventLines) == 0 || len(ventLocation.ventLines) <= y {
				ventLocation.ventLines = append(ventLocation.ventLines,[]string{})
			}
			for x:= ventCoords[k].x[0]; x <= ventCoords[k].x[1]; x ++ {
				for len(ventLocation.ventLines) == 0 || len(ventLocation.ventLines[y]) <= x {
					ventLocation.ventLines[y] = append(ventLocation.ventLines[y],".")
				}
				//var idx int
				//if x != 0 {
				//	idx = x -1
				//}

				//fmt.Println(idx )
				y1, y2 := ventCoords[k].y[0], ventCoords[k].y[1]
				x1, x2 := ventCoords[k].x[0], ventCoords[k].x[1]
				if y1 == y2 || x1 == x2{
					xValueString := ventLocation.ventLines[y][x]
					if xValueString == "." {
						ventLocation.ventLines[y][x] = "1"
						continue
					}
					xValueInt, err := strconv.Atoi(xValueString)
					if err != nil {
						fmt.Println(err)
					}

					if  xValueInt < 5 {
						xValueInt ++
					}
					ventLocation.ventLines[y][x] = strconv.Itoa(xValueInt)
				}

				//var idx int
				//if x != 0 {
				//	idx = x -1
				//}
				//
				//xValueString := ventLocation.ventLines[y][x]
				//if xValueString == "." {
				//	ventLocation.ventLines[y][idx] = "1"
				//	continue
				//}
				//xValueInt, err := strconv.Atoi(xValueString)
				//if err != nil {
				//	fmt.Println(err)
				//}
				//
				//if  xValueInt < 5 {
				//	xValueInt ++
				//}
				//ventLocation.ventLines[y][x] = strconv.Itoa(xValueInt)
			}
		}
	}
	return ventLocation
}

func createVentMap(ventCoords map[int]coords) ventMap {
	keys := make([]int, 0, len(ventCoords))
	for k :=  range ventCoords {
		keys = append(keys, k)
	}

	ventLocation := ventMap{
		ventLines: make([][]string, 0),
	}
	sort.Ints(keys)

	for _, k := range keys {
		sort.Ints(ventCoords[k].x)
		sort.Ints(ventCoords[k].y)
		for y:= ventCoords[k].y[0]; y <= ventCoords[k].y[1]; y ++ {
			for  len(ventLocation.ventLines) == 0 || len(ventLocation.ventLines) <= y {
				ventLocation.ventLines = append(ventLocation.ventLines,[]string{})
			}
			for x:= ventCoords[k].x[0]; x <= ventCoords[k].x[1]; x ++ {
				for len(ventLocation.ventLines) == 0 || len(ventLocation.ventLines[y]) <= x {
					ventLocation.ventLines[y] = append(ventLocation.ventLines[y],".")
				}
				var idx int
				if x != 0 {
					idx = x -1
				}

				xValueString := ventLocation.ventLines[y][x]
				if xValueString == "." {
					ventLocation.ventLines[y][idx] = "1"
					continue
				}
				xValueInt, err := strconv.Atoi(xValueString)
				if err != nil {
					fmt.Println(err)
				}

				if  xValueInt < 5 {
					xValueInt ++
				}
				ventLocation.ventLines[y][x] = strconv.Itoa(xValueInt)
			}
		}
	}
	return ventLocation
}

func (data *coordsData) coordinates(readValue string) error {
	csvString := strings.ReplaceAll(readValue, " -> ", ",")
	splitString := strings.Split(csvString, ",")
	tempCoords := coords{
		x: make([]int, 0),
		y: make([]int, 0),
	}
	for i, strCoord := range splitString {
		coord, err := strconv.Atoi(strCoord)
		if err != nil {
			return err
		}
		if i%2 == 0 {
			tempCoords.x = append(tempCoords.x, coord)
			continue
		}
		tempCoords.y = append(tempCoords.y, coord)
	}

	data.ventCords[data.mapCounter] = tempCoords
	data.mapCounter ++
	return nil
}
