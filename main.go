package main

func main() {
	//initialize a group
	var temp dataCell
	temp.I = 1
	temp.S = "bruh"
	g := groupInit(temp)
	//add a data cell to the group
	var temp2 dataCell
	temp2.I = 2
	temp2.S = "bruh2"
	groupAdd(2, temp2, g)
	//get a data cell from the group
	var temp3 dataCell
	temp3 = groupAt(2, g)
	//print the data cell
	println(temp3.I)
	println(temp3.S)

}
