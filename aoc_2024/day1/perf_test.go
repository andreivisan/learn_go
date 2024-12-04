package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay1P1(b *testing.B) {
    fileName := "input_d1.txt"
    l1, l2 := ReadFile(fileName)
    fmt.Println(Day1_1(l1, l2))
}

func BenchmarkDay1P2(b *testing.B) {
    fileName := "input_d1.txt"
    l1, l2 := ReadFile(fileName)
    fmt.Println(Day1_2(l1, l2))
}
