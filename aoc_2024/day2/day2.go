package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2_1(fileName string) int {
    total := 0
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        fields := strings.Fields(line) 
        if IsLevelSafe(fields) || canRemoveAtMost1(fields) {
            total += 1
        }
    }
    return total
}

func IsLevelSafe(fields []string) bool {
    nums := []int{}
    for _, val := range fields {
        num, _ := strconv.Atoi(val)
        nums = append(nums, num)
        lenNums := len(nums)
        if lenNums > 1 {
            if !isSorted(nums) || absInt(nums[lenNums-1]-nums[lenNums-2]) > 3 {
                return false
            }  
        }
    }
    return true
}

func canRemoveAtMost1(fields []string) bool {
    for i:=0; i < len(fields); i++ {
        fieldsCopy := append([]string{}, fields[:i]...)
        fieldsCopy = append(fieldsCopy, fields[i+1:]...)
        if IsLevelSafe(fieldsCopy) {
            return true
        }
    } 
    return false
}

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func isSorted(arr []int) bool {
    if len(arr) <= 1 {
        return true
    }
    ascending := true
    descending := true
    for i:=1; i<len(arr); i++ {
        if arr[i] < arr[i-1] {
            ascending = false
        } else if arr[i] > arr[i-1] {
            descending = false
        } else {
            return false
        }
    }
    return ascending || descending
}

func main() {
    fileName := "input_d2.txt"
    fmt.Println(Day2_1(fileName))
}
