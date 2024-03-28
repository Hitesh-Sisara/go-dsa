package main

import "fmt"

func main() {
	fmt.Println(search("hitesh", "hite"))
}

func search(haystack, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	for i := 0; i < len(haystack)-len(needle); i++ {

		match := true

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}
