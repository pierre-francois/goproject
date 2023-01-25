package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	mat1 := []byte("1 2 3\n4 5 6\n7 8 9")
	mat2 := byteToInt(mat1)
	fmt.Printf("mat1 : %T\n", mat1)
	fmt.Printf("mat2 : %T\n", mat2)
}
func byteToInt(byteSlice []byte) [][]int {
	//byteSlice := []byte("1 2 3\n4 5 6\n7 8 9")

	// Split the byte slice into rows using the newline character
	rows := bytes.Split(byteSlice, []byte("\n"))

	// Create an empty 2D slice to store the data
	data := make([][]string, len(rows))

	// Iterate over the rows and split each one into columns using the comma character
	for i, row := range rows {
		data[i] = strings.Split(string(row), " ")
	}

	int2D := make([][]int, len(data))
	for i := range int2D {
		int2D[i] = make([]int, len(data[i]))
	}

	// Convert 2D string to 2D int
	for i := range data {
		for j := range data[i] {
			intVal, _ := strconv.Atoi(data[i][j])
			int2D[i][j] = intVal
		}
	}

	//fmt.Println(int2D)
	// Output: [[1 2 3] [4 5 6]]
	return int2D
}
