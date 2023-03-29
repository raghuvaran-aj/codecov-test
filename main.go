package main

import (
	"coverageTesting/sum"
	"fmt"
)

func main() {
	fmt.Println("Main started...")
	n, _ := sum.AddIntegers(1, 2)
	fmt.Println(n)
}
