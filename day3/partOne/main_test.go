package main

import (
	"fmt"
	"testing"
)


func Test_EnergyOutput(T *testing.T) {
input := map[int][]int{
	0: {0,0,0,0,0},
	1: {1,1,1,1,1},
}

	calcEnergy(input)
}


func Test_LifeSupport(T *testing.T) {
	splitInput := map[int][]int{
		0: {0,0,0,0},
		1: {1,1,0,1},
		2: {1,0,0,1},
		3: {0,0,0,0},
	}

	data := []string{
		"01100",
		"01010",
		"00000",
		"00100",
	}

	v, err := calcLifeSupport("high",data, splitInput )
	fmt.Println(v, err)
}