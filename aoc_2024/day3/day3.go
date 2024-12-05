package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3_1(fileName string) int {
    content := getFileContent(fileName)
    return computeResult(content)
}

func Day3_2(fileName string) int {
    content := getFileContent(fileName)
    sum := 0
    enabled := true
    re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
    matches := re.FindAllStringSubmatch(content, -1)
    for _, match := range matches {
        instruction := match[0]
        switch instruction {    
            case "do()":
                enabled = true
            case "don't()":
                enabled = false
            default:
                if enabled {
                    val1, _ := strconv.Atoi(match[1]) 
                    val2, _ := strconv.Atoi(match[2]) 
                    sum += val1 * val2 
                }
        }
    }
    return sum
}

func getFileContent(fileName string) string {
    data, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}

func computeResult(content string) int {
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
    fmt.Println(Day3_2("input_d3.txt"))
}
