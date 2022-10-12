package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	word,long := ChooseWord()
	fmt.Println(word)
	fmt.Println(long)
}

func ChooseWord() (string,int){
	name := os.Args[1]
	body, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	list := []string{}
	hold := ""
	for _, m := range string(body) {
		if m != 10 {
			hold = hold + string(m)
		} else {
			if hold != "" {
				list = append(list, hold)
				hold = ""
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	list = append(list, hold)
	lent := rand.Intn(len(list))
	return list[lent],len(list[lent])-1
}
// func randletters() {
// 	for 0 len mot/2-1
// rand mot 
// mot[rand] tets[rand] == mot[rand]
// }
