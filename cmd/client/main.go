package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("Connected to server")
	//creating a message feature + error handling
	message := []byte("Hello, server to!")
	_, err = conn.Write(message)
	if err != nil {
		log.Println(err)
		return
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Response server: ", response)

}
