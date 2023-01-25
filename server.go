package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net" // for socket resources
	"os"  // for OS resources
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

const (
	Buffer    = 1024 //This constant can be anything from 1 to 65495, because the TCP package can only contain up to 65495 bytes of payload. It will define how big the chunks are of the file that we will send in bytes.
	host      = "localhost"
	port      = "8000"
	protocole = "tcp"
)

//RAJOUTER FUNC GESTION ERREUR

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
func intToByte(Mat *[][]int) []byte {
	MAT := *Mat
	var buf bytes.Buffer
	// Boucler sur chaque élément du slice et les encoder en bytes
	for _, row := range MAT {
		for _, val := range row {
			binary.Write(&buf, binary.LittleEndian, val)
		}
	}
	// Récupérer le tableau de bytes final
	byteArray := buf.Bytes()
	return byteArray
}

func traiterRequete(con net.Conn) {
	buffer := make([]byte, 32) //taille du buffer a set en fonction de la taille de la donnée reçue
	_, err := con.Read(buffer) //lis données reçues et les stocke dans le buffer (ce sont des bytes)
	if err != nil {
		log.Fatal(err)
	}
	mat1 := byteToInt(buffer)
	var pmat1 *[][]int
	pmat1 = &mat1

	buffer2 := make([]byte, 32) //taille du buffer a set en fonction de la taille de la donnée reçue
	a, err := con.Read(buffer2) //lis données reçues et les stocke dans le buffer (ce sont des bytes)
	if err != nil {
		log.Fatal(err)
		print(a)
	}
	mat2 := byteToInt(buffer2)
	var pmat2 *[][]int
	pmat2 = &mat2
	var matResult [][]int
	matResult = prodMat(pmat1, pmat2)
	pmatResult := &matResult
	byte := intToByte(pmatResult)
	fmt.Printf("Taille des donées envoyées : %v\n", len(byte))

	_, err = con.Write(byte) //envoi des données
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	// close conn
	//con.Close()
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

func prodMat(Mata, Matb *[][]int) [][]int {

	MatA := *Mata
	MatB := *Matb
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

	go func(Mata, Matb *[][]int) { //on push des jobs de structure sendJob dans le channel jobChan
		for i := 0; i < ligne; i++ {
			for j := 0; j < col; j++ {
				job := sendJob{Mata, Matb, i, j}
				jobChan <- job
			}
		}
	}(Mata, Matb)

	for b := 0; b < 16; b++ { //on vide le channel de resultat sous forme de structure endJob
		resultJob := <-resultChan
		resultat[resultJob.x][resultJob.y] = resultJob.z

	}
	//close(resultChan) !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//close(jobChan) !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return (resultat)
}
func main() {
	listener, err := net.Listen(protocole, host+":"+port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// interet ???
	defer listener.Close()

	print("Serveur démarré, en attente de connexion.\n")
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go traiterRequete(con) // Action a executer une go routine par requete
	}
}
