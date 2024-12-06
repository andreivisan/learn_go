package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules struct {
    First int
    Second int
}

func Day5_1(fileName string) int {
    firstPart, secondPart := splitFile(fileName)
    rules := parseRules(firstPart)
    sum := 0
    updatesScanner := bufio.NewScanner(strings.NewReader(secondPart))
    for updatesScanner.Scan() {
        line := updatesScanner.Text()
        updates := strings.Split(line, ",")
        mid, ok := validUpdate(rules, updates)
        if ok {
            sum += mid
        }
    }
    return sum
}

func splitFile(fileName string) (string, string) {
    file, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }
    parts := strings.Split(string(file), "\n\n")
    return parts[0], parts[1]
}

func parseRules(firstPart string) (rules []Rules) {
    rulesScanner := bufio.NewScanner(strings.NewReader(firstPart))
    for rulesScanner.Scan() {
        line := rulesScanner.Text()
        rulesLine := strings.Split(line, "|")
        val1, _ := strconv.Atoi(rulesLine[0])
        val2, _ := strconv.Atoi(rulesLine[1])
        rules = append(rules, Rules{val1, val2})
    }
    return rules
}  

func validUpdate(rules []Rules, updates []string) (int, bool) {
    updatesDir := make(map[int]int)
    for idx, val := range(updates) {
        valInt, _ := strconv.Atoi(string(val))
        updatesDir[valInt] = idx
    }
    for _, rule := range rules {
        firstRuleIdx, firstRuleOk := updatesDir[rule.First]
        secondRuleIdx, secondRuleOk := updatesDir[rule.Second]
        if firstRuleOk && secondRuleOk && firstRuleIdx > secondRuleIdx {
            return 0, false
        } 
    }
    mid, _ := strconv.Atoi(updates[len(updates) / 2])
    return mid, true

}

func main() {
    fmt.Println(Day5_1("input_d5.txt"))
}
