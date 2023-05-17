package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func gtail(filename string) error {
	const bufferSize = 70 //byte
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	//https://cs.opensource.google/go/go/+/master:src/io/fs/fs.go;l=151?q=FileInfo&sq=&ss=go

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	/*
		Size() int64        // length in bytes for regular files; system-dependent for others
		here filesize is offset which is the number of bytes to move the file pointer from its current position. For example, if you call file.Seek(10, 0), the file pointer will be moved 10 bytes from the beginning of the file.
		here whence = 2
		0: The file offset is measured from the beginning of the file.
		1: The file offset is measured from the current position of the file pointer.
		2: The file offset is measured from the end of the file.
	*/
	fileSize := fileInfo.Size()
	offsetpos, err := file.Seek(fileSize, 2)
	fmt.Println("first offset position:")
	fmt.Println(offsetpos)
	if err != nil {
		return err
	}
	reader := bufio.NewReaderSize(file, bufferSize)
	//	defaultBufSize = 4096 byte
	for {
		_, err = file.Seek(fileSize, 2)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}
		}
		fmt.Println(offsetpos)
		fmt.Print(line) //print line if it is not empty
		newFileSize, err := file.Stat()
		fmt.Println("newfilesize: " + strconv.FormatInt(newFileSize.Size(), 10))
		if err != nil {
			return err
		}

		if newFileSize.Size() != fileSize {
			fileSize = newFileSize.Size()
			continue // Continue to the next iteration if new lines are appended
		}
		// Sleep for small interval 50ms before reading again
		time.Sleep(900 * time.Millisecond)
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
