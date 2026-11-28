package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("Connected to server")

	reader := bufio.NewReader(os.Stdin)
	for {

		data, err := reader.ReadString('\n') //returns a string along with a line break character
		if err != nil {
			log.Println("Input error", err)
			return
		}

		//command for end work of client
		if strings.TrimSpace(data) == "/exit" {
			fmt.Println("Client disconnecting") // message withou logs
			break
		}

		message := []byte(data)
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
}
