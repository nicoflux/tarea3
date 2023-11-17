package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Matrix struct {
	data [][]string
	rows int
	cols int
	mtx  sync.Mutex
}

func loadMatrix(filename string) *Matrix {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rows, cols int
	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		data = append(data, row)
		rows++
		cols = len(row)
	}

	return &Matrix{data: data, rows: rows, cols: cols}
}

func (m *Matrix) printMatrix() {
	for _, row := range m.data {
		fmt.Println(strings.Join(row, " "))
	}
}

func (m *Matrix) updateMatrix(x, y int, color string, processID int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	// Verificar si la posición ya está pintada con el mismo color
	if m.data[x][y] == color {
		return
	}

	originalColor := m.data[x][y]
	m.floodFill(x, y, color, originalColor)

	// Imprimir el mensaje de salida
	fmt.Printf("Niño %d pintó desde (%d, %d) con color %s.\n", processID, x, y, color)
}

// Función floodFill para matrices de cadenas
func (m *Matrix) floodFill(x, y int, newColor, oldColor string) {
	rows, cols := len(m.data), len(m.data[0])

	// Función interna para llenado recursivo
	var fill func(int, int)
	fill = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || m.data[x][y] != oldColor {
			return
		}

		m.data[x][y] = newColor

		// Llenar recursivamente las posiciones vecinas
		fill(x+1, y)
		fill(x-1, y)
		fill(x, y+1)
		fill(x, y-1)
	}

	// Llamar a la función de llenado desde la posición inicial
	fill(x, y)
}

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Uso: ./main <cantidad_procesos> <filas> <columnas> <matriz.txt> <turnos.txt>")
		os.Exit(1)
	}

	numProcesses, _ := strconv.Atoi(os.Args[1])
	//numRows, _ := strconv.Atoi(os.Args[2])
	//numCols, _ := strconv.Atoi(os.Args[3])
	matrixFileName := os.Args[4]
	turnsFileName := os.Args[5]

	matrix := loadMatrix(matrixFileName)

	var wg sync.WaitGroup
	done := make(chan struct{})

	processTurn := func(processID int) {
		defer wg.Done()

		turnsFile, err := os.Open(turnsFileName)
		if err != nil {
			panic(err)
		}
		defer turnsFile.Close()

		scanner := bufio.NewScanner(turnsFile)
		for scanner.Scan() {
			turn := strings.Fields(scanner.Text())
			x, _ := strconv.Atoi(turn[0])
			y, _ := strconv.Atoi(turn[1])
			color := turn[2]

			matrix.updateMatrix(x, y, color, processID)
		}
	}

	for i := 1; i <= numProcesses; i++ {
		wg.Add(1)
		go processTurn(i)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
	matrix.printMatrix()
}
