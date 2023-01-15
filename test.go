//produit matriciel en go quu marche
package main

type Matrice [][]int

func prodMat(mat1, mat2 Matrice) Matrice {
	col := len(mat2[0])
	ligne := len(mat1)
	resultat := make(Matrice, ligne)

	for i := range resultat {
		resultat[i] = make([]int, col)
	}
	for i := 0; i < ligne; i++ {
		for j := 0; j < col; j++ {
			for k := 0; k < len(mat1); k++ {
				resultat[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
		//}(i)
		//}
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
