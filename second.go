package main

import (
	"fmt"
	"sort"
)

// Creates a map to store the frequency of each character in the string

func rearrangeString(s string) string {
	frequency := make(map[rune]int)

	for _, char := range s {
		frequency[char]++
	}

	// creates a slice and it is sorted based on the character frequency in descending order

	var sortedChars []rune
	for ch := range frequency {
		sortedChars = append(sortedChars, ch)
	}

	sort.Slice(sortedChars, func(i, j int) bool {
		return frequency[sortedChars[i]] > frequency[sortedChars[j]]
	})

	if frequency[sortedChars[0]] > (len(s)+1)/2 {
		return ""
	} //

	result := make([]rune, len(s))
	i := 0
	for _, ch := range sortedChars {
		for j := 0; j < frequency[ch]; j++ {
			if i >= len(s) {
				i = 1
			}
			result[i] = ch
			i += 2
		}
	}

	return string(result)
}

func main() {

	fmt.Println(rearrangeString("aab"))  // output: aba
	fmt.Println(rearrangeString("aaab")) //output: “ “

}
