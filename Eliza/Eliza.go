package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var introductions = []string{
	"Hello, How are you feeling today?",
	"How do you do. Are you seeking help today?",
	"Please tell me what's been bothering you.",
	"Is something troubling you?",
}

var userGoodbyes = []string{
	"bye", "goodbye", "farewell",
}

var goodbyes = []string{
	"Farewell. It was lovely speaking with you.",
	"Thank you for talking with me today.",
	"Thank you, that will be $150. Have a good day!",
	"Goodbye. This was nice, hopefully we do it again sometime.",
	"Goodbye. I'm looking forward to our next session.",
	"Well.. I guess time is up, call back anytime!",
	"Maybe we could discuss this over more in our next session? Goodbye.",
	"Ciao",
}

var reflexivesAfter = map[string]string{
	"you":    "me",
	"me":     "yourself",
	"myself": "yourself",
	"my":     "your",
	"mine":   "yours",
}

var mapMain = map[string][]string{
	"i need ": {"Why do you need %s?",
		"Would it really help you to get %s?",
		"Are you sure you need %s?"},
}

func main() {
	n := randomNumber(len(introductions))
	fmt.Printf("\nEliza:\t" + introductions[n] + "\nYou:\t")
	input()
}

func randomNumber(lenA int) int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(lenA)
	return n
}

func splitString(s string, where string) [250]string {
	var newS [250]string
	var word string
	var firstWord bool = false
	j := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == where {
			for ; string(s[i]) == where; i++ {
			}
			if firstWord {
				newS[j] = word
				j++
				word = ""
			}
		}
		word = word + string(s[i])
		//fmt.Println("word", word)
		firstWord = true
	}
	newS[j] = word
	//fmt.Println("newS", newS)
	return newS
}

func newInput(input string) string {
	var newInput string
	var firstWord bool = false
	for i := 0; i < len(input); i++ {
		if !firstWord {
			for ; string(input[i]) == " "; i++ {
			}
			firstWord = true
		}
		newInput = newInput + string(input[i])
	}
	//fmt.Println("newInput", newInput)
	return newInput
}

func input() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\r\n", "")
		input = strings.ToLower(input)
		response(input)
	}
}

func response(input string) {
	input = newInput(input)
	s := splitString(input, " ")
	//fmt.Println("s", s)
	var s2, key string
	//var j int = 0
	for _, s1 := range s {
		for _, k := range userGoodbyes {
			if s1 == k {
				fmt.Printf("Eliza:\t")
				n := randomNumber(len(goodbyes))
				fmt.Printf(goodbyes[n] + "\n")
				os.Exit(0)
			}
		}
	}
	for _, s1 := range s {
		s2 = s2 + s1 + " "
		//fmt.Println("s2", s2)
		_, ok := mapMain[s2]
		if ok {
			key = s2
			break
		}
	}
	var r string
	var reflexiveCheck bool = false
	for i, s1 := range s {
		_, ok := reflexivesAfter[s1]
		if ok && i > 0 {
			r = reflexivesAfter[s1]
			reflexiveCheck = true
		}
	}
	if !reflexiveCheck {
		r = input[len(key):]
	}
	// fmt.Println("reflexiveCheck", reflexiveCheck)
	// fmt.Println("key", key)
	// fmt.Println("len(key)", len(key))
	// fmt.Println("r", r)
	n := randomNumber(len(mapMain[key]))
	fmt.Printf("Eliza:\t"+mapMain[key][n], r)
	fmt.Printf("\nYou:\t")
}
