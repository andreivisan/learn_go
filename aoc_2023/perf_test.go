package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay1P1(b *testing.B) {
    fileName := "day11.txt"
    fmt.Println(Day1p1Sol(fileName))
}
