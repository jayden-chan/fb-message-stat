package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func write(t thread) {
	fileName := fmt.Sprintf("./assets/template.html")
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	stringValue := string(byteValue)

	stringValue = strings.Replace(stringValue, "__TITLE__", "Conversation with "+t.Title, -1)
	messages := strconv.Itoa(len(t.Messages))
	words := strconv.Itoa(t.words)
	fmt.Println(messages, words)
	stringValue = strings.Replace(stringValue, "__MESS__", messages, -1)
	stringValue = strings.Replace(stringValue, "__WRDS__", words, -1)

	fileTitle := strings.Replace(t.Title, " ", "", -1) + ".html"
	err = ioutil.WriteFile("./out/"+fileTitle, []byte(stringValue), 0644)
	if err != nil {
		panic(err)
	}
}
