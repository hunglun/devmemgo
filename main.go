package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

const (
	devicePath = "/dev/mem"
	dataSize   = 4 // Number of bytes to read/write
)

func main() {
	// Ensure a memory address is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a memory address as a command-line argument")
	}

	// Parse the memory address from the command-line argument
	address, err := strconv.ParseUint(os.Args[1], 0, 64)
	if err != nil {
		log.Fatalf("Failed to parse memory address: %v", err)
	}

	// Open /dev/mem device file
	file, err := os.OpenFile(devicePath, os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		log.Fatalf("Failed to open device file: %v", err)
	}
	defer file.Close()

	// Seek to the desired memory address
	_, err = file.Seek(int64(address), 0)
	if err != nil {
		log.Fatalf("Failed to seek to address: %v", err)
	}

	// Read data from the memory address
	data := make([]byte, dataSize)
	_, err = file.Read(data)
	if err != nil {
		log.Fatalf("Failed to read from address: %v", err)
	}

	// Convert the byte slice to the appropriate data type (e.g., uint32)
	readValue := *(*uint32)(unsafe.Pointer(&data[0]))
	fmt.Printf("Read value: 0x%x\n", readValue)

	// Modify the read value (e.g., increment it)
	writeValue := readValue + 1

	// Convert the modified value to a byte slice
	writeData := (*[dataSize]byte)(unsafe.Pointer(&writeValue))[:]

	// Seek back to the memory address
	_, err = file.Seek(int64(address), 0)
	if err != nil {
		log.Fatalf("Failed to seek to address: %v", err)
	}

	// Write the modified data to the memory address
	_, err = file.Write(writeData)
	if err != nil {
		log.Fatalf("Failed to write to address: %v", err)
	}

	fmt.Println("Write operation completed successfully!")
}
