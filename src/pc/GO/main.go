/**
* main.go
* Author: Jonas S.
* Date  : 15/09/2023
* Brief : This program manages the serial communication with the board
* 	   	  and control the individual volumes set by the user.
**/

// https://learn.microsoft.com/en-us/windows/win32/coreaudio/core-audio-interfaces

package main

import (
	"fmt"
	//"log"
	//"bytes"
	//"strings"
	"time"

	//"github.com/albenik/go-serial"

	// Windows COM library
	// Process management
	// Windows Core Audio

	ole "github.com/go-ole/go-ole"
	ps "github.com/mitchellh/go-ps"
	wca "github.com/moutend/go-wca/pkg/wca"
)

func main() {


	// WTF is this?
    ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
    defer ole.CoUninitialize()

	// ! Create a session manager
	manager, err := wca.CreateAudioSessionManager2()
    if err != nil {
        fmt.Printf("Error creating audio session manager: %v\n", err)
        return
    }
    defer manager.Release()


	// ? Get the default audio session control
	// This is the audio session that is currently playing
	// This is used to get the current volume of a specific application
	control, err := manager.GetDefaultAudioSessionControl(wca.DeviceRoleMultimedia)
    if err != nil {
        fmt.Printf("Error getting default audio session control: %v\n", err)
        return
    }
    defer control.Release()

	volume_app := "Discord.exe"

	// ? Find the process with the given name
	volume_process, err := findProcessByName(volume_app)
    if err != nil {
        fmt.Printf("Error finding process: %v\n", err)
        return
    }
	
	// ? Get the audio session control of the target application
	target_control, err := manager.GetAudioSessionControl(volume_process.(ps.Process).Pid(), 0)
	if err != nil {
		fmt.Printf("Error getting audio session control: %v\n", err)
		return
	}
	defer target_control.Release()

	// ? Get the volume of the target application
	volume, err := target_control.SimpleAudioVolume()
	if err != nil {
		fmt.Printf("Error getting simple audio volume: %v\n", err)
		return
	}
	defer volume.Release()

	// ? Get the current volume of the target application
	current_volume, err := volume.GetMasterVolume()
	if err != nil {
		fmt.Printf("Error getting master volume: %v\n", err)
		return
	}
	fmt.Printf("Current volume: %v\n", current_volume)

	for {
		
		new_volume := current_volume + 0.1
		err = target_control.SetMasterVolume(new_volume)
		if err != nil {
			fmt.Printf("Error setting new volume: %v\n", err)
			return
		}
		fmt.Printf("New volume for %s: %f\n", volume_app, new_volume)
		time.Sleep(1 * time.Second)
	}

	

	//// ! LIST AVAILABLE PORTS
	//ports, err := serial.GetPortsList()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if len(ports) == 0 {
	//	log.Fatal("No serial ports found!")
	//}

	//for _, port := range ports {
	//	fmt.Printf("Found port: %v\n", port)
	//}


	//// ! OPEN SERIAL PORT
	//mode := &serial.Mode{
	//	BaudRate: 115200,
	//}
	//port, err := serial.Open(ports[0], mode)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer port.Close()


	//// ! WAIT FOR PING FROM BOARD
	//receivedData := make([]byte, 0)
	//buf := make([]byte, 10)
	//for {

	//	n, err := port.Read(buf)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	// Append the received data to the buffer
	//	receivedData = append(receivedData, buf[:n]...)
	//	fmt.Println("---")
	//	fmt.Printf("%s\n", receivedData)
	//	fmt.Println("---")

	//	// Check if "PING\n" is present in the received data
	//	if bytes.Contains(receivedData, []byte("PING")) {
	//		fmt.Println("Received PING")
	//		break
	//	}
	//}

	//// ! SEND PONG TO BOARD
	//_, err = port.Write([]byte("PONG\n\n\r"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connection established")

	//// ! WAIT FOR VOLUME COMMANDS
	//// The board will send volume commands in the following format:
	//// REx:{delta_x}
	//// Where x is the Rotary encoder index and {delta_x} is the change in volume
	//buf = make([]byte, 256)
	//var lineBuffer string

	//for {
	//	n, err := port.Read(buf)
	//	if err != nil {
	//		log.Fatalf("port.Read: %v", err)
	//	}

	//	data := string(buf[:n])
	//	lineBuffer += data

	//	if strings.Contains(lineBuffer, "\n") {
	//		lines := strings.Split(lineBuffer, "\n")

	//		// ! Process each line
	//		for _, line := range lines {
				
	//			// Check if the line matches the expected format "REx:{delta_x}"
	//			if strings.HasPrefix(line, "RE") && strings.Contains(line, ":") {
	//				parts := strings.Split(line, ":")
	//				if len(parts) == 2 {
	//					re_index := parts[0]
	//					delta := parts[1]
	//					fmt.Printf("%s -> %s\n", re_index, delta)						
	//				}
	//			}
	//		}

	//		lineBuffer = ""
	//	}
	//}
	
}


// Find a process by name
func findProcessByName(processName string) (*ps.Process, error) {
    processes, err := ps.Processes()
    if err != nil {
        return nil, err
    }

    for _, process := range processes {
        if process.Executable() == processName {
            return &process, nil
        }
    }

    return nil, fmt.Errorf("Process not found: %s", processName)
}

// Find an audio session by process ID
func findSessionByProcessID(control *wca.AudioSessionControl2, processID int) (*wca.AudioSessionControl2, error) {
    sessions, err := control.GetSessionEnumerator()
    if err != nil {
        return nil, err
    }

    defer sessions.Release()

    for _, session := range sessions {
        if session.ProcessID() == uint(processID) {
            return session, nil
        }
    }

    return nil, fmt.Errorf("Session not found for process ID: %d", processID)
}

