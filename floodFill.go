package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func floodFill(matrix [][]rune, x int, y int, NewColor rune, OldColor rune) {
	rows := len(matrix)
	col := len(matrix[0])
	if x < 0 || x >= rows || y < 0 || y >= col {
		return
	}
	if matrix[x][y] != OldColor {
		return
	}

	matrix[x][y] = NewColor

	floodFill(matrix, x+1, y, NewColor, OldColor)
	floodFill(matrix, x-1, y, NewColor, OldColor)
	floodFill(matrix, x, y+1, NewColor, OldColor)
	floodFill(matrix, x, y-1, NewColor, OldColor)
	floodFill(matrix, x+1, y+1, NewColor, OldColor)
	floodFill(matrix, x-1, y-1, NewColor, OldColor)
	floodFill(matrix, x+1, y-1, NewColor, OldColor)
	floodFill(matrix, x-1, y+1, NewColor, OldColor)
}

func main() {

	data, err := os.ReadFile("matrix.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		currentLine := strings.TrimSpace(line)
		currentLine = strings.Join(strings.Fields(currentLine), "")
		lines[i] = strings.TrimSpace(currentLine)
	}

	var matrix [][]rune
	for _, line := range lines {
		row := []rune(line)
		matrix = append(matrix, row)
	}

	fmt.Println("Original matrix:")
	for _, row := range matrix {
		fmt.Println(string(row))
	}

	// Read flood fill data from file
	fillData, err := os.ReadFile("turnos.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	turnosLines := strings.Split(string(fillData), "\n")

	for _, line := range turnosLines {
		data := strings.Split(line, " ")
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		newColor := rune(data[2][0])
		floodFill(matrix, x, y, newColor, matrix[x][y])
	}

	fmt.Println("Matrix after flood fill:")
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}
