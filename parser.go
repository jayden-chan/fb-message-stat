package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type message struct {
	Name    string `json:"sender_name"`
	Time    int    `json:"timestamp_ms"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type thread struct {
	Messages     []message `json:"messages"`
	Title        string    `json:"title"`
	Type         string    `json:"thread_type"`
	Participants []string  `json:"participants"`

	self         string
	distribution map[string]int
}

func (t *thread) computeSelf() {
	for _, m := range t.Messages {
		if !stringContains(t.Participants, m.Name) {
			t.self = m.Name
			return
		}
	}
}

func (t *thread) computeDistribution() {
	for _, m := range t.Messages {
		t.distribution[m.Name]++
	}
}

func initializeThread(path string) (t *thread) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &t)
	if err != nil {
		panic(err)
	}

	t.computeSelf()
	t.distribution = make(map[string]int)
	t.computeDistribution()

	return
}
