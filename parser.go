package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type message struct {
	Name    string `json:"sender_name"`
	Time    int64  `json:"timestamp_ms"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type thread struct {
	Messages     []message `json:"messages"`
	Title        string    `json:"title"`
	Type         string    `json:"thread_type"`
	Participants []string  `json:"participants"`

	messageDist map[string]int
	words       int
	wordDist    map[string]int
}

func (t *thread) computeDistribution() {
	for _, m := range t.Messages {
		t.messageDist[m.Name]++
	}
}

func (t *thread) computeWordCount() {
	for _, m := range t.Messages {
		words := strings.Fields(m.Content)
		t.words += len(words)
		t.wordDist[m.Name] += len(words)
	}
}

func initializeThread(path string) (t *thread) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &t)
	if err != nil {
		panic(err)
	}

	t.messageDist = make(map[string]int)
	t.wordDist = make(map[string]int)
	t.computeDistribution()
	t.computeWordCount()

	return
}
