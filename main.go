package main

import "fmt"

func main() {
	fmt.Println(compare([]int{5, 8, 6, 9}, []int{5, 8, 6, 9, 13}))
}

func stack(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
