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
	PORT = "9090"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	fmt.Println("Trying to connect to server...")
	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected !")

	filetosend, _ := os.Open("matrice.txt")
	_, errr := io.Copy(conn, filetosend)
	if errr != nil {
		log.Fatal(errr)
		return
	}

	/* fi, _ := filetosend.Stat()
	if sentData != fi.Size() {
		fmt.Println("Erreur, je n'ai pas envoyé le bon nombre de bytes")
		return
	}
	fmt.Print("Bytes sent: ", sentData, "\n") */

	buffer := make([]byte, 4096)   //taille du buffer qui reçoit les données
	data, err := conn.Read(buffer) //lit données reçues et les stocke dans le buffer (cbytes)
	if err != nil {
		log.Fatal(err)
		print("erreur")
	}

	file, err := os.Create("resultFile.txt") //crée le fichier où l'on va stocker le résultat
	if err != nil {
		log.Fatal(err)
		return
	}

	n1, err := file.Write(buffer[:data]) // on parcourt le buffer jusqu'à data, taille des données reçues du serveu (le reste est inutile), puis on écrit ces données dans le fichier créé
	fmt.Printf("Received %d bytes\n", n1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Results were written in resultFile.txt")
}
