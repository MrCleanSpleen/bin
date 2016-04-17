package main

import (
	"fmt"
	//c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"log"
	"regexp"
	"strings"
)

func capstring(input string) string {
	words := strings.Fields(input)

	for index, word := range words {
		words[index] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
func remove(str, chr string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	name := i.Ask("Please put in your name: ")
	if strings.ContainsAny(name, "_") {
		name1 := remove(name, "")
		name2 := strings.Replace(name1, "_", " ", -1)
		name3 := capstring(name2)
		name4 := stripchars(name3, " ")
		fmt.Println(name4)
	} else if strings.ContainsAny(name, "-") {
		name1 := remove(name, "")
		name2 := strings.Replace(name1, "-", " ", -1)
		name3 := capstring(name2)
		name4 := stripchars(name3, " ")
		fmt.Println(name4)
	} else {
		fmt.Println(remove(name, ""))
	}
}
