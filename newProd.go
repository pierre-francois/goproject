//produit matriciel en go quu marche
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrice [][]int

type sendJob struct {
	pointA *Matrice
	pointB *Matrice
	i, j   int
}

type endJob struct {
	x, y, z int
}

func calcul(jobChan chan sendJob, resultChan chan endJob) {
	job := <-jobChan
	var res int = 0
	//for k := 0; k < 16; k++ { //changer len(matA)
	//res += (*(job.pointA))[job.i][k] * (*(job.pointB))[k][job.j]
	//}
	result := new(endJob)
	result.x = job.i
	result.y = job.j
	result.z = res
	resultChan <- *result

}

func prodMat(MatA, MatB Matrice) Matrice {

	ligne := len(MatA)
	col := len(MatB[0])
	var resultat Matrice

	jobChan := make(chan sendJob, 40)
	resultChan := make(chan endJob, 40)

	for b := 0; b < 5; b++ {
		go calcul(jobChan, resultChan)
	}

	go func(MatA, MatB Matrice) {
		for i := 0; i < ligne; i++ {
			for j := 0; j < col; j++ {
				job := new(sendJob)
				job.i = i
				job.j = j
				job.pointA = (&MatA)
				job.pointB = (&MatB)
				jobChan <- *job
			}
		}
	}(MatA, MatB)

	for b := 0; b < 16; b++ {
		resultJob := <-resultChan
		resultat[resultJob.x][resultJob.y] = resultJob.z
	}
	close(resultChan)
	close(jobChan)
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
	//var matFile1 string = "matrice1.txt"
	//var matFile2 string = "matrice2.txt"
	var mat1 = Matrice{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	var mat2 = Matrice{{4, 4, 4, 4}, {3, 3, 3, 3}, {2, 2, 2, 2}, {1, 1, 1, 1}}
	//var mat1 Matrice = readMatrice(matFile1)
	//var mat2 Matrice = readMatrice(matFile2)
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
