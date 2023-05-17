package main

import (
	"bufio"
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
			// File truncated, seek to the beginning
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
				log.Println(scanner.Text())
			}

			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}

			// Update the file size
			fileSize = newFileSize
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("gtail <path to file>")
	}

	filename := os.Args[1]

	gtail(filename)
}
