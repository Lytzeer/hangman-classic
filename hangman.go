package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
<<<<<<< HEAD
	"time"
=======
>>>>>>> 1a4a1ca1ba7db1ba4dd0621b6b10e9d494f453a0
)

func main() {
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
<<<<<<< HEAD
	rand.Seed(time.Now().UnixNano())
	list = append(list, hold)
	lent := rand.Intn(len(list))
	fmt.Println(list[lent])
}
=======
	list = append(list, hold)
	lent := rand.Intn(len(list))
	fmt.Println(list[lent])
}
>>>>>>> 1a4a1ca1ba7db1ba4dd0621b6b10e9d494f453a0
