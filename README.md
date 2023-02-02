# goproject

###### _Nicolas BLOCH [nicolas.bloch@insa-lyon.fr] Aurélien GIRIN [aurelien.girin@insa-lyon.fr] Mody Sory SOW [mody.sow@insa-lyon.fr]_

### Project Description

This teamwork was carried out within the framework of the ELP project (___E___*cosystèmes des* ___L___*angages de* ___P___*rogrammation*, programming 
language ecosystems) taught in the telecommunication department of INSA Lyon. 


The aim of this project is to learn the basics of 3 programming languages([JS](https://github.com/jesuisjayus/jsproject "jsproject"),Go & [ELM](https://github.com/jesuisjayus/elmproject "elmproject")) and to produce a functional program.

As far as Go project is concerned, we had to create program that runs a matrix product using goroutines. Version 1 files only contain the product matrix code. For the version 2, we used a TCP connection between a client that sends the matrix, stored in a local file, to a server that computes the product, and returns the matrix result to the client. Each connection is handled by the server with a goroutine

### How to run and use the program 

0. First, make sure golang is installed on your device (if not, check this [page](https://go.dev/doc/install))
1. Download the code and unzip it
2. Go to Version 2 folder and open two terminals
3. First type 'go run server.go' in a terminal and then 'go run client.go' in the other one
4. You should be able to see the matrix result of the product in the resultFile.txt file
