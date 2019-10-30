package helper

import "fmt"

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

func Hello() (i int) {
	fmt.Println("Hello, World!")
	i = 100
	return i
}

/*func CreateCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}*/

func Add(c *Collection, name string) error {
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

func Prev(c *Collection) *Record {
	c.now = c.now.previous
	fmt.Printf("Previous Node: %s", c.now.node)
	return c.now
}
