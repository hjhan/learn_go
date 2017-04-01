package main

import (
	"fmt"
	"math"
	"strings"
)

func named_fun_as_goroutine(s string) {
	fmt.Println(s)

}

const PI = 3.1415926

var (
	a = 5
	b = 10
	c = 15
)

func square(x *float64) {
	*x = *x * *x
}
func main() {
	/**
	Strings
	*/
	str := "HI, I'M UPPER CASE"
	lower := strings.ToLower(str)
	fmt.Println(lower)

	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	sentence := "I'm a sentence made up of	words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	//split on more than just the space, but all whitespace chars
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

	/**
	Pointer
	*/
	f := 1.59990090909090
	square(&f)
	fmt.Println(f)

	/**
	Arrays
	*/
	var arr [5]int
	arr[4] = 100
	fmt.Println(arr)

	/**
	Slices
	*/
	//var s []float64
	s := make([]float64, 5)
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	fmt.Println(s)

	/**
	Maps
	*/
	m := make(map[string]int) //map have to be initialized before use
	m["key"] = 10
	m["k2"] = 99
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "k2")
	fmt.Println(len(m))

	go named_fun_as_goroutine("foobar")

	go func(x int) {
		//func body goes here

	}(42)

	fmt.Println(PI)
	fmt.Println(a, b, c)

	var x string
	x = "Mo Bike"

	fmt.Println(x)
	y := "Hello OFO"
	fmt.Println(y)

	fmt.Printf("hello %s\n", "world!")
	fmt.Printf("Sqrt(9)=%f", math.Sqrt(9.0))
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
}
