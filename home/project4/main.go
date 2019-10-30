package main

import (
	"fmt"
	"project4/helper"
)

type Record struct {
	node     string
	previous *Record
	next     *Record
}

type Collection struct {
	name string
	head *Record
	tail *Record
	now  *Record
}

var MyCollection = &Collection{}

func CreateCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}

func (c *Collection) Add(name string) error {
	n := &Record{
		node: name,
	}
	if c.head == nil {
		c.head = n
	} else {
		currentNode := c.tail
		currentNode.next = n
		n.previous = c.tail
	}
	c.tail = n
	fmt.Print("Add New Node   --> ", n.node, "\n")
	return nil
}

/*func (c *Collection) GetIndex(index int) *Record {
	currentNode := c.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	fmt.Print("Get Index ", index, ": Result is: ", currentNode.node, "\n\n")
	return currentNode
}

func (c *Collection) RemoveIndex(index int) error {
	i := 0
	currentNode := c.head
	//lastNode := c.head
	fmt.Print("Remove Index: ", index, "\n")
	if index == 0 {
		c.head = c.head.next
		return nil
	}

	for currentNode.next != nil {
		i++
		if i == index {
			delNode := currentNode.next
			//fmt.Print("+_+_+_+_+_+_+_+_+", currentNode.next, currentNode.next.next, "\n")
			currentNode.next = currentNode.next.next
			delNode.previous = currentNode
			//fmt.Print("\n+_+_+_+_+_+_+_+_+", delNode, currentNode, "\n")
			//fmt.Print("\n+_+_+_+_+_+_+_+_+", c.head, c.tail.previous, "\n")
			if delNode.next == nil {
				c.tail = c.tail.previous
			}
			return nil
		}
		//lastNode = lastNode.next
		currentNode = currentNode.next
	}
	return nil
}

func (c *Collection) Print() error {
	currentNode := c.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	fmt.Println()
	return nil
}

func (c *Collection) PrintLine() error {
	currentNode := c.head
	split := " --> "
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Print(currentNode.node)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Print(split, currentNode.node)
	}
	fmt.Print("\n\n")
	return nil
}

func (c *Collection) Lenght() (i int) {
	currentNode := c.head
	for currentNode != nil {
		i++
		currentNode = currentNode.next
	}
	return i
}

func (c *Collection) First() *Record {
	return c.head
}

func (c *Collection) Last() *Record {
	return c.tail
}

func (c *Collection) Value() *Record {
	c.now = c.head
	fmt.Printf("Value Node: %s", c.now.node)
	return c.now
}

func (c *Collection) Next() *Record {
	c.now = c.now.next
	fmt.Printf("Next Node: %s", c.now.node)
	return c.now
}*/

func main() {
	MyCollectionName := "Collection One"
	MyCollection := &Collection{}
	fmt.Print("\nCreated Collection --> ", MyCollectionName, "\n")
	fmt.Println(MyCollection)
	MyCollection.Add("Zero")
	MyCollection.Add("qwe")
	/*MyCollection.Add("One")
	MyCollection.Add("Two")
	MyCollection.Add("Three")
	MyCollection.Add("Four")
	/*	fmt.Println()
		//MyCollection.Print()
		MyCollection.PrintLine() //print

		MyCollection.GetIndex(4)
		MyCollection.RemoveIndex(2)

		MyCollection.PrintLine() //print
		fmt.Print("Lenght of Collection is ", MyCollection.Lenght(), " Nodes \n\n")

		//start
		fmt.Print("     --> ", MyCollection.Value(), "\n\n")

		//fmt.Printf("%v, %T", MyCollection.now.previous, MyCollection.now.next)

		//next steps
		fmt.Print("     --> ", MyCollection.Next(), "\n")
		fmt.Print("     --> ", MyCollection.Next(), "\n")
		//prev steps
		//fmt.Print("     --> ", helper.Prev(MyCollection), "\n")
		//fmt.Print("     --> ", helper.Prev(MyCollection), "\n")
		//fmt.Print("     --> ", helper.Prev(MyCollection), "\n")

		fmt.Println()
		MyCollection.Add("Five")
		fmt.Print()
		MyCollection.RemoveIndex(2)
		MyCollection.PrintLine()
		MyCollection.RemoveIndex(3)
		MyCollection.PrintLine()
		fmt.Println("First Node: ", *MyCollection.First())
		fmt.Println("Last Node: ", *MyCollection.Last())
	*/
	fmt.Println(helper.Hello())

	//MyCollection = reverseRecurrsive(MyCollection)

	//fmt.Printf("Now playing: %s\n", MyCollection.now.node)
}
