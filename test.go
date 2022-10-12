package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	//positions := []string{}
	//hangman := []string{}
	bod, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	list2 := []string{}
	hold2 := ""
	for _, d := range string(bod) {
		if d != 10 {
			hold2 = hold2 + string(d)
		} else {
			if hold2 != "" {
				list2 = append(list2, hold2)
				hold2 = ""
			}
		}
	}
	for i:= 7; i<14; i++{
		fmt.Println(list2[i])

	}
}

