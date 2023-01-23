package main

import (
	"bufio"
	"fmt"
	"log"
	"net" // for socket resources
	"os"  // for OS resources
	"strconv"
	"strings"
)

const (
	Buffer    = 1024 //This constant can be anything from 1 to 65495, because the TCP package can only contain up to 65495 bytes of payload. It will define how big the chunks are of the file that we will send in bytes.
	host      = "localhost"
	port      = "8000"
	protocole = "tcp"
)

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
		go ReceiveFile(con) // Action a executer une go routine par requete
	}
}

func ReceiveFile(con net.Conn) {
	fmt.Println("connection avec un client !\n")

	defer con.Close()

	buffer := make([]byte, 1024)
	n, err := con.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := string(buffer[:n])

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for {
		n, err := con.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n == 0 {
			break
		}
		file.Write(buffer[:n])
	}

	fmt.Println("File received:", fileName)
	result := readMatrice(fileName)
	for i := 0; i < 4; i++ {
		for j := 0; i < 4; i++ {
			print(result[i][j])
			print(" ")
		}
		println()
	}

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
