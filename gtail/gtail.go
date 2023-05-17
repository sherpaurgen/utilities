package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func gtail(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the initial file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()

	for {
		// Check if the file size has changed
		newFileInfo, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		newFileSize := newFileInfo.Size()

		if newFileSize < fileSize {
			/* File truncated, seek to the beginning
			 whence: 0 means relative to the origin of the file, 1 means relative to the current offset,
			and 2 means relative to the end
			*/
			_, err = file.Seek(0, 0)
			if err != nil {
				log.Fatal(err)
			}
		} else if newFileSize > fileSize {
			// New lines appended, read and print them
			_, err = file.Seek(fileSize, 0)
			if err != nil {
				log.Fatal(err)
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			/* https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/bufio/scan.go;l=29
			this shows that if scanner gets EOF a <nil> is set to scanner.Err()
			*/
			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}

			// Update the file size
			fileSize = newFileSize
		}

		time.Sleep(900 * time.Millisecond)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("gtail <path to file>")
	}

	filename := os.Args[1]

	gtail(filename)
}
