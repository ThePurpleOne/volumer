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

	"github.com/karalabe/hid"
)

func main(){

    vid := 0x1234 // VENDOR ID (VID)
    pid := 0x6969 // PRODUCT ID (PID)

	devices, err := hid.Enumerate(vid, pid)
    if err != nil {
        log.Fatal(err)
    }
    defer hid.Close()
	
    if len(devices) == 0 {
        fmt.Println("No matching HID devices found")
        return
    }

    // Open the first matching device
    device, err := devices[0].Open()
    if err != nil {
        log.Fatal(err)
    }
    defer device.Close()

    // Read HID reports
    buf := make([]byte, 2)
    for {

        n, err := device.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        // Process the received HID report
        if n > 0 {
            fmt.Printf("Received HID report: %v\n", buf[:n])
        }

    }
}
