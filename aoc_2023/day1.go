package main

import (
	"fmt"
	"unicode"
)

func sumDigits(line string) int {
    sum := 0
    runes := []rune(line)
    p1  := 0
    p2  := len(runes) - 1
    for p1 < len(runes) {
        if unicode.IsDigit(runes[p1]) {
            sum += int(runes[p1] - '0')
            break
        } else {
            p1 += 1
        }
    }
    for p2 > 0 {
        if unicode.IsDigit(runes[p2]) {
            sum += int(runes[p2] - '0')
            break
        } else {
            p2 -= 1
        }
    }
    return sum
}

func main() {
    line := "pqr34tu8vwx"
    fmt.Println(sumDigits(line))
}
