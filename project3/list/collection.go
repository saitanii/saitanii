package list

import "fmt"

type Collection struct {
	name string
	head *Record
	tail *Record
	now  *Record
}

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

func (c *Collection) GetIndex(index int) *Record {
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
}

func (c *Collection) Prev() *Record {
	c.now = c.now.previous
	fmt.Printf("Previous Node: %s", c.now.node)
	return c.now
}
