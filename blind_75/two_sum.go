package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        if j, ok := m[target-v]; ok {
            return []int{i, j}
        }
        m[v] = i
    }
    return []int{}
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}