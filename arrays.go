package main

import "fmt"

type group struct {
	head *node
	tail *node
}

type node struct {
	//group cell
	cell dataCell
	//next node
	next *node
	//index of the node
	index int
}
type dataCell struct {
	I int
	B bool
	S string
	F float64
	C rune
}

func initGroup(data ...dataCell) group {
	//create a new group
	var newGroup group
	//create a new node
	newNode := new(node)
	newNode.cell = data[0]
	newNode.next = nil
	//set the head and tail
	newGroup.head = newNode
	newGroup.tail = newNode
	//add the rest of the nodes
	if len(data) > 1 {
		temp := newNode // temp is a pointer to the first node
		for i := 1; i < len(data); i++ {
			newNode = new(node)    // create a new node
			newNode.cell = data[i] // set the cell
			newNode.next = nil     // set the next node to nil
			temp.next = newNode    // set the next node of the previous node to the new node
			temp = newNode         // set the temp to the new node
		}
	}
	return newGroup
}
func addNode(data dataCell, group group) {
	//create a new node
	var newNode node
	newNode.cell = data
	newNode.next = nil
	//add the new node to the tail
	group.tail.next = &newNode
	//update the tail
	group.tail = &newNode
}
func delNode(index int, group group) {
	//find the node
	var currentNode node
	currentNode = *group.head
	for currentNode.index != index {
		currentNode = *currentNode.next
	}
	//delete the node
	currentNode = *currentNode.next
}
func printGroup(group group) {
	//print the group
	var currentNode node
	currentNode = *group.head
	for currentNode.next != nil {
		fmt.Println(currentNode.cell)
		currentNode = *currentNode.next
	}
	fmt.Println(currentNode.cell)
}
func readNode(index int, group group) dataCell {
	//find the node
	var currentNode node
	currentNode = *group.head
	for currentNode.index != index {
		currentNode = *currentNode.next
	}
	//return the cell
	return currentNode.cell
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
