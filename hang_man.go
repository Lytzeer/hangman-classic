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
	list2 := PosHangman()
	fmt.Println("Good Luck, you have 10 attempts.")
	word, long := ChooseWord()
	fmt.Println(word)
	var nouvmot string
	for i := 0; i < len(word)-1; i++ {
		nouvmot += string(word[i])
	}
	mot, attempts := InitGame(word)
	mott := ShowWord(mot)
	fmt.Println(mott)
	Play(attempts, nouvmot, mot, long, list2)
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

func Play(attempts int, word string, mottab []string, long int, list2 []string) {
	count := 0
	var present bool
	var letter string
	for word != TabtoStr(mottab) {
		if attempts == 0 {
			for i := len(list2) - 8; i < len(list2)-1; i++ {
				fmt.Println(list2[i])
			}
			fmt.Println("")
			fmt.Println("Game over! The correct word was", word)
			return
		} else {
			present = false
			fmt.Print("Choose: ")
			fmt.Scan(&letter)
			if len(letter) > 1 {
				if letter == word {
					fmt.Println("Congrats !")
					return
				} else {
					attempts--
					count += 7
				}
			}
			for i := 0; i < len(word); i++ {
				if string(word[i]) == letter {
					mottab[i] = letter
					present = true
				}
			}
		}
		if !present {
			attempts--
			if attempts >= 1 {
				fmt.Println("Not present in the word, ", attempts, " attempts remaining")
				for num := count; num < count+7; num++ {
					fmt.Println(list2[num])
				}
				fmt.Println("")
			}
			count += 7
		}
		fmt.Println(TabtoStr(mottab))
	}
	fmt.Println("Congrats !")
}

func ShowWord(word []string) string {
	var motstr string
	for _, ch := range word {
		motstr += " " + ch
	}
	return motstr
}

func TabtoStr(word []string) string {
	str := ""
	for _, ch := range word {
		str += ch
	}
	return str
}

func PosHangman() []string {
	list2 := []string{}
	bod, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
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
	list2 = append(list2, hold2)
	return list2
}
