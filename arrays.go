package main

import "errors"

func groupInit(a ...dataCell) *group {
	//initialize a group with the given data cells
	g := new(group)
	var tempLast *group
	i := 1
	for i = 1; i < len(a)+1; i++ {
		temp := new(group)
		if i == len(a) {
			temp.next = nil
		} else if i != 1 {
			tempLast.next = temp
		} else {
			g.next = temp
		}
		temp.first = g
		temp.index = i
		temp.value = a[i-1]
		tempLast = temp

	}
	tempBruh := tempLast.first
	for i = 1; i < len(a)+1; i++ {
		tempBruh.last = tempLast
	}
	return g
}
func groupAt(index int, bruh *group) dataCell {
	temp := bruh
	for temp.next != nil {
		if temp.index == index {
			return temp.value
		}
		temp = temp.next
	}
	var nope dataCell
	return nope

}
func groupAdd(index int, data dataCell, bruh *group) error {
	temp := bruh
	var err error = nil
	for {
		if temp.index == index {
			err = errors.New("Index already exists")
			return err
		}
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	g := new(group)
	temp.next = g
	g.index = index
	g.value = data
	g.first = bruh.first
	g.last = g
	g.next = nil
	temp = bruh.first
	for temp.next != nil {
		temp.last = g
		temp = temp.next
	}
	return err
}

type group struct {
	//group cell
	first *group
	last  *group
	index int
	value dataCell
	next  *group
}
type dataCell struct {
	I int
	B bool
	S string
	F float64
	C rune
}

/*func groupSort(b []group) []group {
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
*/
