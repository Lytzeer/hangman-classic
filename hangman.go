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
	solution, long := ChooseWord()
	fmt.Println(solution)
	fmt.Println(long)
	//mottab, attempts := InitGame(solution)
	//fmt.Println(mottab)
	//fmt.Println(attempts)
	//gameword := ShowWord(mottab)
	// fmt.Println(gameword)
	// fmt.Println(len(gameword))
}

func ChooseWord() (string, int) {
	name := os.Args[1]
	rep := []string{}
	mot := ""
	content, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	} else {
		for _, ch := range content {
			if ch == 13 {
				rep = append(rep, mot)
				mot = ""
			} else if string(ch) != " " {
				mot = mot + string(ch)
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	rep = append(rep, mot)
	lent := rand.Intn(len(rep))
	word := ""
	for i := 0; i < len(rep[lent]); i++ {
		fmt.Println(string(rep[lent][i]))
	}
	return word, len(rep[lent])
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

func TabtoStr(word []string) string {
	var str string
	for _, ch := range word {
		str += ch
	}
	return str
}

func PlayGame(solution string, mottab []string) {

}
