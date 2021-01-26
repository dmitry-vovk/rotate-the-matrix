package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the matrixRotation function below.
func matrixRotation(matrix [][]int32, r int32) {
	rows := len(matrix)
	cols := len(matrix[0])
	// find out the number of layers
	layers := rows / 2
	if min := cols / 2; min < layers {
		layers = min
	}
	// go layer by layer clockwise
	for l := 0; l < layers; l++ {
		var queue []int32
		// top row
		for col := l; col < cols-l; col++ {
			queue = append(queue, matrix[l][col])
		}
		// right col
		for row := l + 1; row < rows-l-1; row++ {
			queue = append(queue, matrix[row][cols-l-1])
		}
		// bottom row
		for col := l; col < cols-l; col++ {
			queue = append(queue, matrix[rows-l-1][cols-col-1])
		}
		// left col
		for row := l + 1; row < rows-l-1; row++ {
			queue = append(queue, matrix[rows-row-1][l])
		}
		lr := int(r)
		if ql := len(queue); ql < int(r) {
			lr = int(r) % ql
		}
		// Rotate the layer by lr
		queue = append(queue[lr:], queue[:lr]...)
		n := 0
		// top row
		for col := l; col < cols-l; col++ {
			matrix[l][col] = queue[n]
			n++
		}
		// right col
		for row := l + 1; row < rows-l-1; row++ {
			matrix[row][cols-l-1] = queue[n]
			n++
		}
		// bottom row
		for col := l; col < cols-l; col++ {
			matrix[rows-l-1][cols-col-1] = queue[n]
			n++
		}
		// left col
		for row := l + 1; row < rows-l-1; row++ {
			matrix[rows-row-1][l] = queue[n]
			n++
		}
	}
	// output rotated matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println("")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	mnr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(mnr[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mnr[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(mnr[2], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var matrix [][]int32
	for i := 0; i < int(m); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(n) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	matrixRotation(matrix, r)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
