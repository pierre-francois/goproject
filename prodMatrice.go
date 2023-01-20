//produit matriciel en go quu marche
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type sendJob struct {
	pointA *[][]int
	pointB *[][]int
	i, j   int
}

type endJob struct {
	x, y, z int
}

var matFile1 string = "matrice1.txt"
var matFile2 string = "matrice2.txt"

func calcul(jobChan chan sendJob, resultChan chan endJob) {
	for true {
		job := <-jobChan
		var res int
		var m1 = *(job.pointA)
		var m2 = *(job.pointB)
		for k := 0; k < len(m2); k++ {
			res += m1[job.i][k] * m2[k][job.j]
		}
		result := endJob{job.i, job.j, res}
		resultChan <- result
	}
}

func prodMat(MatA, MatB [][]int) [][]int {

	var ligne int = len(MatA)
	var col int = len(MatB[0])

	resultat := make([][]int, ligne)
	for m := range resultat {
		resultat[m] = make([]int, col)
	}

	jobChan := make(chan sendJob, 2)
	resultChan := make(chan endJob, 2)

	for b := 0; b < 4; b++ {
		go calcul(jobChan, resultChan)
	}

	go func(MatA, MatB [][]int) {
		for i := 0; i < ligne; i++ {
			for j := 0; j < col; j++ {
				job := sendJob{&MatA, &MatB, i, j}
				jobChan <- job
			}
		}
	}(MatA, MatB)

	for b := 0; b < 16; b++ {
		resultJob := <-resultChan
		resultat[resultJob.x][resultJob.y] = resultJob.z

	}
	//close(resultChan)
	//close(jobChan)
	return resultat
}

func readMatrice(fileMat string) [][]int {
	file, err := os.Open(fileMat)
	if err != nil {
		log.Fatal(err) // affiche l'erreur et quitte la fonction
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // Créez un nouveau scanner pour lire le fichier
	var matrice [][]int               // Déclarez une variable pour stocker la matrice

	for scanner.Scan() { // Pour chaque ligne du fichier...

		line := scanner.Text()               // Obtenez la ligne courante
		elements := strings.Split(line, " ") // Découpez la ligne en une liste d'éléments

		var row []int // Pour chaque élément de la ligne...
		for _, element := range elements {

			num, err := strconv.Atoi(element) // Convertir l'élément en entier
			if err != nil {
				log.Fatal(err)
			}

			// Ajoutez l'élément converti à la ligne
			row = append(row, num)
		}

		// Ajoutez la ligne à la matrice
		matrice = append(matrice, row)
	}
	return (matrice)
}

func verifMat(MatA, MatB [][]int) bool {
	if len(MatA[0]) != len(MatB) {
		return false
	}
	return true
}

func main() {
	//var mat1 = [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	//var mat2 = [][]int{{4, 4, 4, 4}, {3, 3, 3, 3}, {2, 2, 2, 2}, {1, 1, 1, 1}}

	var mat1 [][]int = readMatrice(matFile1)
	var mat2 [][]int = readMatrice(matFile2)
	var ligne int = len(mat1)
	var col int = len(mat2)

	if !verifMat(mat1, mat2) {
		print("Erreur : impossible de multriplier les matrices.\nVeuillez verifier leurs dimensions\n")
	} else {
		var result = prodMat(mat1, mat2)
		print("Le produit de la matrice A par la matrice B est :\n")
		for i := 0; i < ligne; i++ {
			for j := 0; j < col; j++ {
				print((result[i][j]))
				if result[i][j] < 10 {
					print("  ")
				} else {
					print(" ")
				}
			}
			println()
		}
	}
}
