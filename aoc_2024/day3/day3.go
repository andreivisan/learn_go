package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3_1(fileName string) int {
    data, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }
    content := string(data)
    re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
    matches := re.FindAllStringSubmatch(content, -1)
    sum := 0
    for _, match := range matches {
        val1, _ := strconv.Atoi(match[1]) 
        val2, _ := strconv.Atoi(match[2]) 
        sum += val1 * val2    
    }
    return sum
}

func main() {
    fmt.Println(Day3_1("input_d3.txt"))
}
