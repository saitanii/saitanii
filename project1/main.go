package main

import (
	"fmt"
	"project1/helper"
)

//"hello/helper"

func main() {
	s := helper.Arr(10, 1)
	helper.Arrprint(s)
	helper.Change(s, 7)
	fmt.Println("___________")
	helper.Arrprint(s)

}
