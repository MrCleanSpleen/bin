package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"strings"
)

func CapString(input string) string {
	words := strings.Fields(input)

	for index, word := range words {
		words[index] = strings.Title(word)
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
	if strings.ContainsAny(name, "_") {
		strings.Replace(name, "_", " ", -1)
		fmt.Println(StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;-'+="))
		fmt.Println(c.CL)
		CapString(name)
		fmt.Println(name)
	} else if strings.ContainsAny(name, "-") {
		strings.Replace(name, "-", " ", -1)
		fmt.Println(StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;_'+="))
		fmt.Println(c.CL)
		CapString(name)
		fmt.Println(name)
	} else {
		fmt.Println(StripChars(name, "~!@#$%^&*()<>.,?/|}{[]:;_- +='"))
	}
}
