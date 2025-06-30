package main

import (
	"fmt"
)

func longestString(strings ...string) string {
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := longestString("cat", "apple", "dog", "blue", "red")
	fmt.Println("Самая длинная строка:", result)
}
