package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(fileName string) (l1 []int, l2 []int) {
    var list1 []int
    var list2 []int
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        fields := strings.Fields(line)

        num1, _ := strconv.Atoi(fields[0])
        list1 = append(list1, num1)
        num2, _ := strconv.Atoi(fields[1])
        list2 = append(list2, num2)
    }
    sort.Ints(list1)
    sort.Ints(list2)     
    return list1, list2 
}

func mapOccurences(l1 []int, l2 []int) map[int]int {
    lenL1 := len(l1)
    countOcc := make(map[int]int)
    i, j := 0, 0
    for i < lenL1 {
        x := l1[i]
        if _, exists := countOcc[x]; !exists {
            lenL2 := len(l2)
            for j < lenL2 && l2[j] < x {
                j++
            }
            count := 0
            tempJ := j
            for tempJ < lenL2 && l2[tempJ] == x {
                count++
                tempJ++
            }
            countOcc[x] = count
            j = tempJ
        }
        i++
    }
    return countOcc
}

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func Day1_1(l1 []int, l2 []int) int {
    sum := 0
    lenL1 := len(l1)
    for i:=0; i<lenL1; i++ {
        sum += absInt(l1[i] - l2[i])
    }
    return sum
}

func Day1_2(l1 []int, l2 []int) int {
    sum := 0
    countOcc := mapOccurences(l1, l2)
    for _, num := range(l1) {
        sum += num * countOcc[num]
    }
    return sum
}

func main() {
    l1, l2 := ReadFile("input_d1.txt")
    fmt.Println(Day1_1(l1, l2))
    fmt.Println(Day1_2(l1, l2))
}
