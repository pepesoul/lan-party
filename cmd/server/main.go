package main

import (
	"fmt"
	"io"
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
		conn, err := listener.Accept() //accepting connection
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

	for {

		n, err := conn.Read(buffer) // n - count of byte which it was possible to read

		if n > 0 {
			message := string(buffer[:n]) // takes only a received bytes using - [:n] and converts them into a string
			fmt.Println("Received: ", message)

			//sending a message to a client

			_, writeErr := conn.Write([]byte(message)) //converts a string into bytes
			if writeErr != nil {
				log.Println(err)
				return
			}
		}

		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
			} else {
				log.Println("Read error", err)
			}
			return
		}

	}
}
