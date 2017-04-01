package main

import "fmt"

const MAX = 1000

func main() {
	sum := 0
	for i := 0; i < MAX; i++ {
		sum += check(i)
	}
	fmt.Printf("Total: %d \n", sum)
}
func check(n int) int {
	if (n%3 == 0) || (n%5 == 0) {
		return n
	}
	return 0
}
