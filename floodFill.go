package main

import "fmt"

// Structure pour représenter une position dans la matrice
type Position struct {
	X int
	Y int
}

func FloodFill(matrix [][]rune, point Position, testColor rune, newChar rune) {
	fmt.Println("got inside FloodFill")

	rows := len(matrix)
	columns := len(matrix[0])
	fmt.Println("rows: ", rows, " columns: ", columns)
	if point.X < 0 || point.X >= rows || point.Y < 0 || point.Y >= columns {
		fmt.Println("Point ", point, " est hors de la matrice")
		return
	}

	if matrix[point.X][point.Y] != testColor {
		fmt.Println("Point ", point, " n'est pas de la bonne couleur")
		return
	}

	matrix[point.X][point.Y] = newChar

	FloodFill(matrix, Position{point.X - 1, point.Y}, testColor, newChar)
	FloodFill(matrix, Position{point.X + 1, point.Y}, testColor, newChar)
	FloodFill(matrix, Position{point.X, point.Y - 1}, testColor, newChar)
	FloodFill(matrix, Position{point.X, point.Y + 1}, testColor, newChar)
	/* 	if point.X > 0 && point.Y > 0 {
	   		FloodFill(matrix, Position{point.X - 1, point.Y - 1}, testColor, newChar)
	   	}
	   	if point.X < columns-1 && point.Y > 0 {
	   		FloodFill(matrix, Position{point.X + 1, point.Y - 1}, testColor, newChar)
	   	}
	   	if point.X > 0 && point.Y < rows-1 {
	   		FloodFill(matrix, Position{point.X - 1, point.Y + 1}, testColor, newChar)
	   	}
	   	if point.X < columns-1 && point.Y < rows-1 {
	   		FloodFill(matrix, Position{point.X + 1, point.Y + 1}, testColor, newChar)
	   	} */
}

func main() {
	// Matrice d'exemple
	matrix := [][]rune{
		{'.', '.', '.', '.'},
		{'.', 'O', 'O', '.'},
		{'.', 'O', 'O', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', 'O', 'O'},
		{'.', '.', 'O', 'O'},
	}

	// Afficher la matrice originale
	fmt.Println("Matrice originale:")
	for _, row := range matrix {
		fmt.Println(string(row))
	}
	var point Position
	point.X = 2
	point.Y = 4
	fmt.Println("calling FloodFill")
	FloodFill(matrix, point, matrix[point.X][point.Y], 'X')

	// Afficher la matrice modifiée
	fmt.Println("Matrice après flood fill:")
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}
