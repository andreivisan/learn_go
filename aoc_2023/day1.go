package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func day1p1Sol(fileName string) int {
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
    fmt.Println(day1p1Sol(fileName))
}
