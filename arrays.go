package main

import "fmt"

type field struct {
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

func initGroup(data ...dataCell) field {
	//create a new group
	var newGroup field
	//create a new node
	/*
		newNode := node{}
		newNode.cell = data[0]
		newNode.next = nil
		newNode.index = 0*/
	newNode := &node{cell: data[0], next: nil, index: 0}
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
			newNode.index = i      // set the index
			temp.next = newNode    // set the next node of the previous node to the new node
			temp = newNode         // set the temp to the new node
		}
	}
	return newGroup
}
func (g *field) add(data dataCell) {
	//create a new node
	var newNode node
	newNode.cell = data
	newNode.next = nil
	//add the new node to the tail
	g.tail.next = &newNode
	//update the tail
	g.tail = &newNode
}
func (g *field) delete(index int) {
	//find the node
	var currentNode node
	currentNode = *g.head
	for currentNode.index != index {
		currentNode = *currentNode.next
	}
	//delete the node
	currentNode = *currentNode.next
}
func printGroup(g field) {
	//print the group
	var currentNode node
	currentNode = *g.head
	for currentNode.next != nil {
		fmt.Println(currentNode.cell)
		currentNode = *currentNode.next
	}
	fmt.Println(currentNode.cell)
}
func (g field) at(index int) dataCell {
	//find the node
	var currentNode node
	currentNode = *g.head
	for currentNode.index != index {
		if currentNode.next == nil {
			return dataCell{}
		}
		currentNode = *currentNode.next
	}
	//return the cell
	return currentNode.cell
}
