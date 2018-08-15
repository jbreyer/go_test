package anagrams

import (
	"strings"
)

func Anagram(word1 string, word2 string) bool{

	if !(len(word1) == len(word2)){
		return false
	}

	for _, letter := range(word1){
		if (strings.Count(word1, string(letter)) != strings.Count(word2, string(letter))){
			return false
		}
	}

	return true
}