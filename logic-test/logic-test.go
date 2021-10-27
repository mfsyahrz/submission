package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(FindAnagram(words))
}

func FindAnagram(words []string) [][]string {

	//create map to store group of anagrams
	wordMap := make(map[string][]string)

	// loop through words
	for idx, word := range words {

		// sort word by whatever order and set as keyword for map
		keyword := []rune(word)
		sort.Slice(keyword, func(i, j int) bool {
			return keyword[i] > keyword[j]
		})

		// append map based on keyword (sorted word)
		wordMap[string(keyword)] = append(wordMap[string(keyword)], words[idx])
	}

	anagramList := [][]string{}

	// loop through map to pass anagram to slice
	for _, wordGroup := range wordMap {
		anagramList = append(anagramList, wordGroup)
	}

	return anagramList
}
