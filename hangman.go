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
	fmt.Println("Good Luck, you have 10 attempts.")
	word, long := ChooseWord()
	fmt.Println(word)
	fmt.Println(long)
	mot, attempts := InitGame(word)
	fmt.Println(mot, attempts)
	mott := ShowWord(mot)
	fmt.Println(mott)
}

func ChooseWord() (string, int) {
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
	return list[lent], len(list[lent]) - 1
}

func InitGame(word string) ([]string, int) {
	mot := []string{}
	for i := 0; i < len(word)-1; i++ {
		mot = append(mot, "_")
	}
	var letterreveal int
	for i := 0; i < (len(word)/2)-1; i++ {
		letterreveal = rand.Intn(len(mot))
		mot[letterreveal] = string(word[letterreveal])
	}
	return mot, 10
}

func ShowWord(word []string) string {
	var motstr string
	for _, ch := range word {
		motstr += " " + ch
	}
	return motstr
}
