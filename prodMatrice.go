//produit matriciel en go quu marche
package main

import "sync"

type Matrice [][]int

func calcul(matA, matB Matrice, canal chan int, wg *sync.WaitGroup, i, j int) {
	defer wg.Done() //revoir doc mais apparement indispensable
	var res int
	for k := 0; k < len(matA); k++ {
		res += matA[i][k] * matB[k][j]
	}
	canal <- res //on renvoi le resultat par l'intermédiaire du canal
}

func prodMat(mat1, mat2 Matrice) Matrice {

	ligne := len(mat1)
	col := len(mat2[0])

	resultat := make(Matrice, ligne, col)
	for i := range resultat {
		resultat[i] = make([]int, col)
	}

	var wg sync.WaitGroup       // on crée un wait group, necessaire d'attendre que tte les go routines finissent avant de close le channel
	canal := make(chan int, 16) //on crée le canal pour renvoyer les valeurs calculées
	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			wg.Add(1)                               //on incrément le waitGroup d'une goroutine
			go calcul(mat1, mat2, canal, &wg, i, j) //on pointe vers le waitgroup créé
		}
	}
	wg.Wait()    //on attend la finalisation de toutes les goroutines
	close(canal) //on ferme le canal

	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			resultat[i][j] = <-canal //on rempli la matrice avec le resultat des go-routines
		}
	}
	return resultat

}

func main() {
	var mat1 = Matrice{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	var mat2 = Matrice{{4, 4, 4, 4}, {3, 3, 3, 3}, {2, 2, 2, 2}, {1, 1, 1, 1}}
	var result Matrice = prodMat(mat1, mat2)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			print((result[i][j]))
			print(" ")
		}
		println()
	}
}
