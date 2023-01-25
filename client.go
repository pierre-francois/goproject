package main

import (
	"fmt"
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
	print("Tentative de connexion...\n")

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	//connexion au socket
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	print("Connexion établie !\n")

	byte1, _ := os.ReadFile("matrice1.txt")
	fmt.Printf("Taille des donées envoyées : %v\n", len(byte1))

	_, err = conn.Write(byte1) //envoi des données
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	byte2, _ := os.ReadFile("matrice2.txt")
	fmt.Printf("Taille des donées envoyées : %v\n", len(byte2))
	_, err = conn.Write(byte2)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	buffer2 := make([]byte, 1024) //taille du buffer a set en fonction de la taille de la donnée reçue
	__, err := conn.Read(buffer2) //lis données reçues et les stocke dans le buffer (ce sont des bytes)
	if err != nil {
		log.Fatal(err)
		print(__)
	}
	print(string(buffer2))

	file, err := os.Create("resultatFinal.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	n1, err := file.Write(buffer2)
	fmt.Printf("wrote %d bytes\n", n1)
	conn.Close()

	/* n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
	}
	if n == 0 {
		conn.Close()
	} */
}
