package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func Day1p1Sol(fileName string) int {
    result := 0
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        result += sumDigits(fileScanner.Text())
    }
    return result
}

func Day1p2Sol(fileName string) int {
    result := 0
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        result += sumDigitsAndLiterals(fileScanner.Text())
    }
    return result
}

func sumDigitsAndLiterals(line string) int {
    digits := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    nums := []int{}
    runes := []rune(line)
    lenRunes := len(runes)
    for i:=0; i < lenRunes; i++ {
        if unicode.IsDigit(runes[i]) {
            nums = append(nums, int(runes[i] - '0'))
        } else {
            for k, v := range digits {
                kRunes := []rune(k)     // Convert the key to a slice of runes
                kLen := len(kRunes)  
                if i + kLen <= lenRunes && string(runes[i:i+kLen]) == k {
                    nums = append(nums, v)
                }
            }
        }
    }
    fmt.Println(nums)
    return nums[0] * 10 + nums[len(nums) - 1]
}

func sumDigits(line string) int {
    sum := 0
    runes := []rune(line)
    p1  := 0
    p2  := len(runes) - 1
    for p1 < len(runes) {
        if unicode.IsDigit(runes[p1]) {
            sum = sum * 10 + int(runes[p1] - '0')
            break
        } else {
            p1 += 1
        }
    }
    for p2 >= 0 {
        if unicode.IsDigit(runes[p2]) {
            sum = sum * 10 + int(runes[p2] - '0')
            break
        } else {
            p2 -= 1
        }
    }
    return sum
}

func main() {
    fileName := "day11.txt"
    fmt.Println(Day1p2Sol(fileName))
}
