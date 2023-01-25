package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var mat1 = [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	mat2 := intToByte(mat1)
	fmt.Printf("mat1: %T\n", mat1)
	fmt.Printf("mat2: %T\n", mat2)
}
func intToByte(Mat [][]int) []byte {
	var buf bytes.Buffer
	// Boucler sur chaque élément du slice et les encoder en bytes
	for _, row := range Mat {
		for _, val := range row {
			binary.Write(&buf, binary.LittleEndian, val)
		}
	}
	// Récupérer le tableau de bytes final
	byteArray := buf.Bytes()
	return byteArray
}
