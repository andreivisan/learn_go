package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day4_1(fileName string) int {
    grid := readGrid(fileName)
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

func Day4_2(fileName string) int {
    grid := readGrid(fileName)
    count := 0
    lenRows := len(grid)
    lenCols := len(grid[0])
    for i:=0; i<lenRows; i++ {
        for j:=0; j<lenCols; j++ {
            if(hasMasX(grid, i, j)){
                count += 1
            }
        }
    }
    return count
}

func hasMasX(grid [][]rune, i, j int) bool {
    if i < 1 || i >= len(grid)-1 || j < 1 || j >= len(grid[0])-1 {
        return false
    }
    if grid[i][j] != rune('A') {
        return false
    }
    posibilities := []string{"MS", "SM"}
    diag1 := string(grid[i-1][j-1]) + string(grid[i+1][j+1])
    diag2 := string(grid[i-1][j+1]) + string(grid[i+1][j-1])
    return containsString(posibilities, diag1) && containsString(posibilities, diag2)
}

func containsString(slice []string, str string) bool {
    for _, item := range slice {
        if item == str {
            return true
        }
    }
    return false
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
    fmt.Println(Day4_2("input_d4.txt"))
}
