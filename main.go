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
	fmt.Println("Messages:", len(thread.Messages))
	fmt.Println("Distribution:", thread.messageDist)
	fmt.Println("By words:", thread.wordDist)
	fmt.Println("Total words:", thread.words)
	fmt.Println("Average words per message:", float64(thread.words)/float64(len(thread.Messages)))
}
