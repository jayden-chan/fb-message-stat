package main

import (
	"flag"
	"fmt"
)

func main() {
	filePath := flag.String("path", "", "path to the JSON file containing the message thread")
	flag.Parse()

	fmt.Println(*filePath)
	thread := initializeThread(*filePath)
	fmt.Println(thread.self)
	fmt.Println(thread.distribution)
}
