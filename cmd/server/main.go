package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080") //creating a tcp-listener
	if err != nil {
		log.Fatal(err) //error output and terminates the program
	}
	defer listener.Close()

	fmt.Println("Server started")

	for {
		conn, err := listener.Accept() //accepting one connection
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client connected")

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	//creating a buffer
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer) // n - count of byte which it was possible to read
	if err != nil {
		log.Println(err) //error output, but only the error message is displayed
		return
	}

	message := string(buffer[:n]) // takes only a received bytes using - [:n] and converts them into a string
	fmt.Println("Received: ", message)

	//sending a message to a client
	_, err = conn.Write([]byte(message)) //converts a string into bytes
	if err != nil {
		log.Println(err)
		return
	}

}
