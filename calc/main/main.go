package main

import (
	"fmt"

	"example.com/operations"
)

func main() {
	fmt.Println(operations.Add(1, 2))
	fmt.Println(operations.Minus(5, 5))
	fmt.Println(operations.Multiply(5, 6))
	fmt.Println(operations.Divide(5, 7))
}

/*
	For local use
	go mod edit -replace (module path)=(location of module)
	go get (module path)
	go mod tidy
*/
