package main

import (
	"fmt"
	"project3/list"
)

func main() {
	myCollectionName := "Collection One"
	myCollection := list.CreateCollection(myCollectionName)
	fmt.Print("\nCreated Collection --> ", myCollectionName, "\n")
	fmt.Println()
	myCollection.Add("Zero")
	myCollection.Add("One")
	myCollection.Add("Two")
	myCollection.Add("Three")
	myCollection.Add("Four")
	fmt.Println()
	//myCollection.Print()
	myCollection.PrintLine() //print

	myCollection.GetIndex(4)
	myCollection.RemoveIndex(2)

	myCollection.PrintLine() //print
	fmt.Print("Lenght of Collection is ", myCollection.Lenght(), " Nodes \n\n")

	//start
	fmt.Print("     --> ", myCollection.Value(), "\n\n")

	//fmt.Printf("%v, %T", myCollection.now.previous, myCollection.now.next)

	//next steps
	fmt.Print("     --> ", myCollection.Next(), "\n")
	fmt.Print("     --> ", myCollection.Next(), "\n")
	//prev steps
	fmt.Print("     --> ", myCollection.Prev(), "\n")
	fmt.Print("     --> ", myCollection.Prev(), "\n")
	fmt.Print("     --> ", myCollection.Prev(), "\n")

	fmt.Println()
	//myCollection.Add("Five")
	fmt.Print()
	myCollection.RemoveIndex(2)
	myCollection.PrintLine()
	myCollection.RemoveIndex(3)
	myCollection.PrintLine()
	fmt.Println("First Node: ", *myCollection.First())
	fmt.Println("Last Node: ", *myCollection.Last())

	//myCollection = reverseRecurrsive(myCollection)

	//fmt.Printf("Now playing: %s\n", myCollection.now.node) */
}
