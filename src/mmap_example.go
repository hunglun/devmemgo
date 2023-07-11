package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	filePath     = "file.txt"
	fileSize     = 1024
	writeContent = 0xaa55
)

func main() {
	// Open the file in read-write mode
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Truncate the file to the desired size
	if err := file.Truncate(fileSize); err != nil {
		log.Fatalf("Failed to truncate file: %v", err)
	}

	// Use mmap to map the file into memory
	data, err := syscall.Mmap(int(file.Fd()), 0, fileSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("Failed to mmap file: %v", err)
	}
	defer syscall.Munmap(data)

	// Write the desired content to the start of the file
	*(*uint16)(unsafe.Pointer(&data[0])) = writeContent

	fmt.Println("Data written successfully!")
}
