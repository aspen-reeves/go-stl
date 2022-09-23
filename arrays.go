package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func groupInit(a ...dataCell) []group {
	array := make([]group, len(a))
	for i := 0; i < len(a); i++ {
		array[i] = group{
			index: i,
			value: a[i],
		}
	}
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	return array
}
func groupAt(index int, a []group) dataCell {
	//group the data cells at index
	for i := 0; i < len(a); i++ {
		if a[i].index == index {
			return a[i].value
		}
	}
	return dataCell{}
}

type group struct {
	//group cell
	index int
	value dataCell
}
type dataCell struct {
	I int
	B bool
	S string
	F float64
	C rune
}

func groupSort(b []group) []group {
	//sort the array using bogo sort
	//convert the array to an array of ints

	intArray := make([]int, len(b)*5)
	j := 0
	for i := 0; i < len(intArray); i += 5 {
		a := groupAt(j, b)
		intArray[i] = a.I
		if a.B {
			intArray[i+1] = 1
		} else {
			intArray[i+1] = 0
		}
		temp, err := strconv.Atoi(a.S)
		if err != nil {
			intArray[i+2] = 0
		} else {
			intArray[i+2] = temp
		}
		intArray[i+3] = int(a.F)
		intArray[i+4] = int(a.C)
		j++
	}
	fmt.Println(intArray)
	intArray = sort(intArray)
	fmt.Println(intArray)
	a := make([]group, len(intArray)/5)
	//convert the array of ints back to an array of groups
	j = 0
	for i := 0; i < len(intArray); i += 5 {
		a[j].index = j
		a[j].value.I = intArray[i]
		if intArray[i+1] == 1 {
			a[j].value.B = true
		} else {
			a[j].value.B = false
		}
		a[j].value.S = strconv.Itoa(intArray[i+2])
		a[j].value.F = float64(intArray[i+3])
		a[j].value.C = rune(intArray[i+4])
		j++
	}

	return a
}

func sort(a []int) []int {
	//sort the array using bogo sort
	rand.Seed(time.Now().UnixNano())
	for !isSorted(a) {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
	return a
}
func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
