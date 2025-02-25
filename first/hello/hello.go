package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Glass())
}

/*
	1. Initialize the project using <go mod init (module path)>
	2. Create a .go file
	3. Run the go file using <go run . / go run (filename).go / go run *.go>
	4. To import an external package you can either:
		Import it directly with <go get <package name>
		or
		Type the package you want to input directly into the file and invoke <go mod tidy> (preferrable method)
*/
