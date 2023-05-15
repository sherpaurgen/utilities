package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func gtail(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	//https://cs.opensource.google/go/go/+/master:src/io/fs/fs.go;l=151?q=FileInfo&sq=&ss=go
	//
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	//	Size() int64        // length in bytes for regular files; system-dependent for others
	fileSize := fileInfo.Size()
	//here filezise is offset which is the number of bytes to move the file pointer from its current position. For example, if you call file.Seek(10, 0), the file pointer will be moved 10 bytes from the beginning of the file.
	// here whence = 2
	//0: The file offset is measured from the beginning of the file.
	//1: The file offset is measured from the current position of the file pointer.
	//2: The file offset is measured from the end of the file.
	_, err = file.Seek(fileSize, 2)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	//	defaultBufSize = 4096 byte
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}
		}

		if line != "" {
			fmt.Print(line) //print line if it is not empty
		}

		// Sleep for small interval 50ms before reading again
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run gtail.go <filename>")
		return
	}

	filename := os.Args[1]

	err := gtail(filename)
	if err != nil {
		log.Fatal(err)
	}
}
