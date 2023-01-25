package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sendJob struct {
	pointA *Matrice
	pointB *Matrice
	i, j   int
}

type endJob struct {
	x, y, z int
}

type Matrice [][]int

var matFile1 string = "matrice1.txt"
var matFile2 string = "matrice2.txt"
var resultProduit string = "resultProd.txt"

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

func prodMat(MatA, MatB Matrice) {

	var ligne int = len(MatA)
	var col int = len(MatB[0])

	resultat := make(Matrice, ligne)
	for m := range resultat {
		resultat[m] = make([]int, col)
	}

	jobChan := make(chan sendJob, 2)
	resultChan := make(chan endJob, 2)

	for b := 0; b < 4; b++ {
		go calcul(jobChan, resultChan)
	}

	go func(MatA, MatB Matrice) { //on push des jobs de structure sendJob dans le channel jobChan
		for i := 0; i < ligne; i++ {
			for j := 0; j < col; j++ {
				job := sendJob{&MatA, &MatB, i, j}
				jobChan <- job
			}
		}
	}(MatA, MatB)

	for b := 0; b < 16; b++ { //on vide le channel de resultat sous forme de structure endJob
		resultJob := <-resultChan
		resultat[resultJob.x][resultJob.y] = resultJob.z

	}
	//close(resultChan)
	//close(jobChan)
	writeMatrice(resultat)
}

func readMatrice(fileMat string) Matrice {
	file, err := os.Open(fileMat)
	if err != nil {
		log.Fatal(err) // affiche l'erreur et quitte la fonction
	}

	scanner := bufio.NewScanner(file) // Crée un nouveau scanner pour lire le fichier
	var matrice Matrice

	for scanner.Scan() { // Pour chaque ligne du fichier...

		line := scanner.Text()               // Obtient la ligne courante
		elements := strings.Split(line, " ") // Découpe la ligne en une liste d'éléments

		var row []int
		for _, element := range elements { // Pour chaque élément de la ligne

			num, err := strconv.Atoi(element) // Convertit l'élément en entier
			if err != nil {
				log.Fatal(err)
			}

			// Ajoute l'élément converti à la ligne
			row = append(row, num)
		}

		// Ajoutezla ligne à la matrice
		matrice = append(matrice, row)
	}
	return (matrice)
}

func verifMat(MatA, MatB Matrice) bool {
	if len(MatA[0]) != len(MatB) {
		return false
	}
	return true
}

func writeMatrice(matrice Matrice) {
	file, err := os.Create(resultProduit)
	if err != nil {
		fmt.Println(err)
		return
	}

	ligne := len(matrice)
	col := len(matrice[0])
	for i := 0; i < ligne; i++ {
		for j := 0; j < col-1; j++ {
			fmt.Fprint(file, matrice[i][j], " ")
		}
		fmt.Fprint(file, matrice[i][col-1]) //si on laisse le dernier element de la ligne dans la boucle, un espacé est écrit après et on arrive pas a lire le fichier
		fmt.Fprintln(file)
	}
}

func afficheMatrice(matrice Matrice) {
	ligne := len(matrice)
	col := len(matrice[0])

	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			print((matrice[i][j]))
			if matrice[i][j] < 10 {
				print("  ")
			} else {
				print(" ")
			}
		}
		println()
	}
}

func main() {
	//var mat1 = Matrice{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	//var mat2 = Matrice{{4, 4, 4, 4}, {3, 3, 3, 3}, {2, 2, 2, 2}, {1, 1, 1, 1}}

	var MatA Matrice = readMatrice(matFile1)
	var MatB Matrice = readMatrice(matFile2)

	if !verifMat(MatA, MatB) {
		print("Erreur : impossible de multriplier les matrices.\nVeuillez verifier leurs dimensions\n")
	} else {
		prodMat(MatA, MatB)
		resultat := readMatrice(resultProduit)
		afficheMatrice((resultat))
	}
}
