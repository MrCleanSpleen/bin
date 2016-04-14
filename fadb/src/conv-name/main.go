package main

import (
	"fmt"
	i "github.com/skilstak/go-input"
	"strings"
)
func CapString(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func main() {
	name := i.Ask("Please put in your name: ")
	if name strings.ContainsAny("_") {
		strings.Replace(name, "_", " ", -1)
		stripchars(name, "~!@#$%^&*() <>.,?/|}{[]:;"))
		CapString(name)
	}
 
}
