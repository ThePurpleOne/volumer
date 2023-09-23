/**
* main.go
* Author: Jonas S.
* Date  : 15/09/2023
* Brief : This program is used to read the HID device data
* 	   	  and control the the individual volumes set by the user.
**/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

func main() {

	var vid uint16 = 0x1234
	var pid uint16 = 0x6969 

	device, err := hid.OpenFirst(vid, pid)
	if err != nil {
		log.Fatalf("Failed to open HID device: %v", err)
	}

	//device, err := hid.Open(vid, pid, "")
	//if err != nil {
	//	log.Fatalf("Failed to open HID device: %v", err)
	//}
	//defer device.Close()

	buf := make([]byte, 64) // Assuming HID reports are 64 bytes

	for {
		n, err := device.Read(buf)
		if err != nil {
			log.Fatalf("Error reading from HID device: %v", err)
		}

		if n > 0 {
			report := buf[:n]
			fmt.Printf("Received HID report: %v\n", report)
		}

		time.Sleep(100 * time.Millisecond) // Adjust the sleep duration as needed
	}
}
