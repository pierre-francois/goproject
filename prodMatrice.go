//produit matriciel en go quu marche
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Matrice [][]int

func calcul(matA, matB Matrice, canal chan int, wg *sync.WaitGroup, i, j int) {
	var res int = 0
	for k := 0; k < len(matA); k++ { //changer len(matA)
		res += matA[i][k] * matB[k][j]
	}
	canal <- res
}

func prodMat(mat1, mat2 Matrice) Matrice {

	ligne := len(mat1)
	col := len(mat2[0])

	var resultat Matrice

	var wg sync.WaitGroup
	canal := make(chan int, 16)
	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				calcul(mat1, mat2, canal, &wg, i, j)
			}(i, j)
		}
	}

	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			resultat[i][j] <- canal
		}
	}
	wg.Wait()
	close(canal)
	return resultat

}

func readMatrice(fileMat string) Matrice {
	file, err := os.Open(fileMat)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // Créez un nouveau scanner pour lire le fichier
	var matrice Matrice               // Déclarez une variable pour stocker la matrice

	for scanner.Scan() { // Pour chaque ligne du fichier...

		line := scanner.Text()               // Obtenez la ligne courante
		elements := strings.Split(line, " ") // Découpez la ligne en une liste d'éléments

		var row []int // Pour chaque élément de la ligne...
		for _, element := range elements {

			num, err := strconv.Atoi(element) // Convertir l'élément en entier
			if err != nil {
				fmt.Println(err)
				//return
			}

			// Ajoutez l'élément converti à la ligne
			row = append(row, num)
		}

		// Ajoutez la ligne à la matrice
		matrice = append(matrice, row)
	}

	// Affichez la matrice
	return (matrice)
}

func main() {
	var matFile1 string = "matrice1.txt"
	var matFile2 string = "matrice2.txt"
	//var mat1 = Matrice{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	//var mat2 = Matrice{{4, 4, 4, 4}, {3, 3, 3, 3}, {2, 2, 2, 2}, {1, 1, 1, 1}}
	var mat1 Matrice = readMatrice(matFile1)
	var mat2 Matrice = readMatrice(matFile2)
	//print(mat1)
	//print(mat2)
	var result Matrice = prodMat(mat1, mat2)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			print((result[i][j]))
			print(" ")
		}
		println()
	}
}