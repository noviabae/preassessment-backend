package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWord := string(runes)

		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func main() {
	words := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(words)
	fmt.Println(result)
}
