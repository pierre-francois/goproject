package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type sendJob struct {
	pointA *[][]int
	pointB *[][]int
	i, j   int
}

var A [][]int
var B [][]int

type endJob struct {
	x, y, z int
}

const (
	HOST = "localhost"
	PORT = "9090"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	print("En attente de conexion...\n")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		con, err := listen.Accept() //accepte la connexion
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go traiterRequete(con) //une go routine par requete client
	}

	listen.Close() //jamais atteint, boucle while juste au-dessus
}

func traiterRequete(con net.Conn) {
	fmt.Println("Connexion réussie avec le client !!!")

	reader := bufio.NewReader(con) //lit le contenu des données envoyées par le client

	fileData, _ := reader.ReadString(';') //récupère le contenu correspondant à la première matrice
	mat1 := fileData[:len(fileData)-1]    //supprime le point virgule
	lines := strings.Split(mat1, "\n")    //sépre les données par lignes
	MatA := make([][]int, len(lines))     //crée le slice où l'on va stocker la matrice
	for i, line := range lines {          //pour chaque ligne obtenue
		values := strings.Split(line, " ") //on sépare les valeurs espcées d'un " "
		MatA[i] = make([]int, len(values))
		for j, value := range values {
			num, _ := strconv.Atoi(value) //on convertit l'élément en entier
			MatA[i][j] = num              //on le copie dans la matrice
		}
	}

	fileData2, _ := reader.ReadString(';')  // on prend le contenu correspondant à la mtrice suivante ...
	mat2 := fileData2[1 : len(fileData2)-1] //on supprime le point virgule et un caractère en trop au déut (pk ????)
	lines2 := strings.Split(mat2, "\n")
	MatB := make([][]int, len(lines2))
	for i, line := range lines2 {
		values2 := strings.Split(line, " ")
		MatB[i] = make([]int, len(values2))
		for j, value := range values2 {
			num, _ := strconv.Atoi(value)
			MatB[i][j] = num
		}
	}

	fmt.Println(MatA)
	fmt.Println(MatB)

	result := prodMat(MatA, MatB) //on calcul le produit matriciel
	toSend := intToString(result) //on convertit en string le resultat

	fmt.Print(toSend)
	con.Write([]byte(toSend)) //on convertit en byte et on envoie
	con.Close()
}
func intToString(arr [][]int) string {
	var str string
	for _, row := range arr {
		for _, item := range row {
			str += fmt.Sprint(item) + " "
		}
		str = strings.TrimRight(str, " ") + "\n"
	}
	return str
}
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

	go func(MatA, MatB [][]int) { //on push des jobs de structure sendJob dans le channel jobChan
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
	return resultat
}
