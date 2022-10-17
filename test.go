package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	save()
}

func save() {

	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err == nil {
		ioutil.WriteFile("save.txt", b, 0644)
	}
	err = json.Unmarshal(b, &m)
	fmt.Println(&m)
	fmt.Println(m.Body)
}

type Message struct {
	Name string
	Body string
	Time int64
}
