package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:18060")
	if err != nil {
		fmt.Println(err)
		return
	}

	var mode string
	fmt.Print("Enter mode (PUB/SUB): ")
	_, err = fmt.Scanln(&mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	if mode == "PUB" {
		fmt.Print("Enter topic: ")
		var topic string
		_, err = fmt.Scanln(&topic)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter message: ")
		var message string
		_, err = fmt.Scanln(&message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Send the input to the server
		_, err = conn.Write([]byte(fmt.Sprintf("PUBL:%s,%s", topic, message)))
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error receiving response:", err)
			return
		}
		response := string(buffer[:n])
		// Process the response
		fmt.Println(response)
		return
	} else if mode == "SUB" {
		fmt.Print("Enter topic: ")
		var topic string
		_, err = fmt.Scanln(&topic)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Send the input to the server
		_, err = conn.Write([]byte(fmt.Sprintf("SUBC:%s", topic)))
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error receiving response:", err)
				return
			}

			response := string(buffer[:n])

			// Process the response
			fmt.Println(response)
		}
	}

	// Close the connection
	conn.Close()
}
