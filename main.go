package main

import "fmt"

func main() {
	//create a new group
	var newGroup group
	//insert data
	newGroup = initGroup(dataCell{0, false, "no", 1.1, 'a'}, dataCell{1, true, "hello", 1.1, 'a'}, dataCell{2, false, "world", 2.2, 'b'}, dataCell{3, true, "hello", 3.3, 'c'}, dataCell{4, false, "world", 4.4, 'd'}, dataCell{5, true, "hello", 5.5, 'e'}, dataCell{6, false, "world", 6.6, 'f'}, dataCell{7, true, "hello", 7.7, 'g'}, dataCell{8, false, "world", 8.8, 'h'}, dataCell{9, true, "hello", 9.9, 'i'}, dataCell{10, false, "world", 10.10, 'j'})

	//print the group
	fmt.Println(readNode(5, newGroup))
}
