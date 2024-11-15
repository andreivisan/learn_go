package main

import "fmt"

func containsDuplicate(nums []int) bool {
    seen := make(map[int]int)

    for _, val := range nums {
        if _, ok := seen[val]; ok {
            return true
        }
        seen[val] += 1
    }

    return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}