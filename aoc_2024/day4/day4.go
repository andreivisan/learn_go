package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day4_1(fileName string) int {
    directions := []struct{dx, dy int}{
        {-1, 0},  //up
        {-1, 1},  //up-right
        {0, 1},   //right
        {1, 1},   //down-right
        {1, 0},   //down
        {1, -1},  //down-left
        {0, -1},  //left
        {-1, -1}, //up-left
    } 
    grid := readGrid(fileName)
    count := 0
    lenRows := len(grid)
    lenCols := len(grid[0])
    for i:=0; i<lenRows; i++ {
        for j:=0; j<lenCols; j++ {
            for _, direction := range directions {
                if isXmas(grid, i, j, direction.dx, direction.dy, "XMAS") {
                    count += 1
                }
            }
        }
    }
    return count
}

func isXmas(grid [][]rune, i, j, dx, dy int, word string) bool {
    lenRows := len(grid)
    lenCols := len(grid[0])
    lenWord := len(word)
    for k:=0; k<lenWord; k++ {
        xAxis := i + k * dx
        yAxis := j + k * dy
        if xAxis < 0 || xAxis >= lenRows || yAxis < 0 || yAxis >= lenCols {
            return false
        } 
        if grid[xAxis][yAxis] != rune(word[k]) {
            return false
        }
    }
    return true
}

func readGrid(fileName string) [][]rune {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var grid [][]rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, []rune(line))
    }
    return grid
}

func main() {
    fmt.Println(Day4_1("input_d4.txt"))
}
