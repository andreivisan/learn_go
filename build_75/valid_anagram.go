package main

import "fmt"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    seen := make(map[rune]int)
    for _, v := range s {
       seen[v] += 1
    }
    for _, v := range t {
        if _, ok := seen[v]; ok {
            seen[v] -= 1
        } else {
            return false
        }
    }
    for _, v := range seen {
        if v != 0 {
            return false
        }
    }
    return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}