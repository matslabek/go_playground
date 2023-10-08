package main

import (
	"golang.org/x/tour/wc";
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		elem, isPresent := wordMap[word]
		_ = elem // Just to avoid "elem declared and not used" err
		if isPresent {
			wordMap[word]++	
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
