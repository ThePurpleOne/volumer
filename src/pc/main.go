/**
* main.go
* Author: Jonas S.
* Date  : 15/09/2023
* Brief : This program manages the serial communication with the board
* 	   	  and control the individual volumes set by the user.
**/

package main

import (
	"fmt"
	"log"
	"bytes"
	"strings"
	//"time"

    "github.com/albenik/go-serial"


)

func main() {


	// ! LIST AVAILABLE PORTS
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}


	// ! OPEN SERIAL PORT
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()


	// ! WAIT FOR PING FROM BOARD
	receivedData := make([]byte, 0)
	buf := make([]byte, 10)
	for {

		n, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// Append the received data to the buffer
		receivedData = append(receivedData, buf[:n]...)
		fmt.Println("---")
		fmt.Printf("%s\n", receivedData)
		fmt.Println("---")

		// Check if "PING\n" is present in the received data
		if bytes.Contains(receivedData, []byte("PING")) {
			fmt.Println("Received PING")
			break
		}
	}

	// ! SEND PONG TO BOARD
	_, err = port.Write([]byte("PONG\n\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established")

	// ! WAIT FOR VOLUME COMMANDS
	// The board will send volume commands in the following format:
	// REx:{delta_x}
	// Where x is the Rotary encoder index and {delta_x} is the change in volume
	buf = make([]byte, 256)
	var lineBuffer string

	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}

		data := string(buf[:n])
		lineBuffer += data

		if strings.Contains(lineBuffer, "\n") {
			lines := strings.Split(lineBuffer, "\n")

			// ! Process each line
			for _, line := range lines {
				
				// Check if the line matches the expected format "REx:{delta_x}"
				if strings.HasPrefix(line, "RE") && strings.Contains(line, ":") {
					parts := strings.Split(line, ":")
					if len(parts) == 2 {
						re_index := parts[0]
						delta := parts[1]
						fmt.Printf("%s -> %s\n", re_index, delta)						
					}
				}
			}

			// Clear the lineBuffer
			lineBuffer = ""
		}
	}
	
}
