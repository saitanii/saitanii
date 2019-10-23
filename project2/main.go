package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	r := make(map[int64]int64)

	for i := 0; i < 1000; i++ {
		r[randInt()+randInt()]++
	}

	fmt.Println(r)
}

func randInt() int64 {
	return rand.Int63n(6) + 1
}
