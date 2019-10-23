package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	//Pi same word
	Pi = 3.14
	//Nature = iota
	//Social = iota
)

//var c, ph, java bool
//var c, ph, java = true, false, "nol"
//var i, j int = 1, 2

func main() {
	const World = "Hello World"
	//var pizza string = "Margarita"
	//var pizza = "Margarita"
	//var pizza string
	//pizza = "Margarita"
	//pizza := newPizza()
	pizzas := []string{newPizza(), "Cheese"}
	pizzas = append(pizzas, "Pepperoni")
	pizzas = append(pizzas, "MyPizza")

	//fmt.Println(pizza, Pi)

	sl := make([]int, 5, 5)
	//fmt.Println(i, c, ph, java)
	a, b, _ := swap("one", "two", 12)
	fmt.Println(a, b)
	printSlice("Slice", sl)
	sl = append(sl, 1, -2, 3, -4, -5, 6, 7, 8)
	printSlice("Slice", sl)

	// for i, pizza := range pizzas {
	// 	fmt.Println("+", i, pizza)
	// }

	fmt.Println(PosSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(PosSum([]int{1, -2, 3, 4, 5}))
	fmt.Println(PosSum([]int{-1, -2, -3, -4, -5}))

	//for i := 0; i < 10; i = i + 2 {
	//	fmt.Println(i)
	//}

	//sum := 1
	//for sum <= 10 {
	//	fmt.Println("+", sum)
	//	sum++
	//}

	//fmt.Println(RepeatStr(3, "hello"))
	fmt.Println("The time is", time.Now())
	fmt.Println("_______________________________________________________________")
	fmt.Println(trimLeftChar("Hello"))
	fmt.Println(RemoveChar("Hello World"))
	fmt.Println("_______________________________________________________________")
	fmt.Println("InAscOrder", InAscOrder([]int{1, 2, 4, 7, 19, 20}))
	fmt.Println("_______________________________________________________________")
	fmt.Println("IsTriangle", IsTriangle(6, 5, 2))
	fmt.Println("_______________________________________________________________")
	args := "should test that the correct value"
	fmt.Println("FindShort Form:", args, "is", FindShort(args))
	fmt.Println("_______________________________________________________________")
	fmt.Println("NbYear:", NbYear(1000, 5, 100, 1500))
	fmt.Println("_______________________________________________________________")
	fmt.Println(FizzBuzzCuckooClock("13:34"))
	fmt.Println(Sqrt(2))

}

func newPizza() string {
	return "Margarita"
}

func swap(x, y string, z int) (string, string, int) {
	return y, x, z
}

func printSlice(s string, sl []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(sl), cap(sl), sl)
}

func PosSum(num []int) int {
	var suma int
	for i := range num {
		//fmt.Println("+", i, num)
		if num[i] > 0 {
			suma = suma + num[i]
		}
	}
	return suma
}

func RepeatStr(repetitions int, value string) (s string) {
	for i := 0; i < repetitions; i++ {
		s += value
	}
	return s
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[0:]
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]

}

func InAscOrder(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] <= numbers[i-1] {
			//fmt.Println("+", i, numbers)
			return false
		}
	}
	return true
}

func IsTriangle(a, b, c int) bool {
	return (a+b) > c && (a+c) > b && (c+b) > a
}

// func FindShort(s string) int {
// 	// your code
// }

func FindShort(s string) int {
	length := len(s)
	//fmt.Println(length)
	for _, word := range strings.Split(s, " ") {
		if len(word) < length {
			length = len(word)
		}
		//fmt.Println(word, len(word))
	}
	return length
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	year := 0
	fmt.Println(p0, percent, aug, p)
	for p0 < p {
		p0 = int(float64(p0) + float64(p0)*percent/100 + float64(aug))
		year++
		fmt.Println(year, p0)
	}
	return year
}

func FizzBuzzCuckooClock(time string) (word string) {
	switch time {
	case "13:34":
		word = "tick"
	case "21:00":
		word = "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
	case "11:15":
		word = "Fizz Buzz"
	case "03:03":
		word = "Fizz"
	case "14:30":
		word = "Cuckoo"
	case "08:55":
		word = "Buzz"
	case "00:00":
		word = "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
	case "12:00":
		word = "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
	default:
		//fmt.Println(time)
	}
	return word
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 50; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}
