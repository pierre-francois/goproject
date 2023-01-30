package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8000"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	errorManager((err))
	fmt.Println("Trying to connect to server...")

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	errorManager((err))

	fmt.Println("Connected !")

	filetosend, err := os.Open("matrice.txt") //lit le fichier contenant les matrices
	errorManager((err))

	_, er := io.Copy(conn, filetosend) //copy le fichier et l'envoie
	errorManager((er))

	buffer := make([]byte, 4096)   //taille du buffer qui reçoit les données
	data, err := conn.Read(buffer) //lit données reçues et les stocke dans le buffer (cbytes)
	errorManager((err))

	file, err := os.Create("resultFile.txt") //crée le fichier où l'on va stocker le résultat
	errorManager((err))

	n1, err := file.Write(buffer[:data]) // on parcourt le buffer jusqu'à data, taille des données reçues du serveu (le reste est inutile), puis on écrit ces données dans le fichier créé
	fmt.Printf("Received %d bytes\n", n1)
	errorManager((err))

	fmt.Println("Results were written in resultFile.txt")

	//lire le fichier contenant le résultat

	/* content, err := ioutil.ReadFile("resultFile.txt")
	errorManager((err))
	fmt.Println(string(content)) */
}

func errorManager(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
