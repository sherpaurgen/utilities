package main

//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"time"
//)
//
//func gtail(filename string) {
//	file, err := os.Open(filename)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	// Get the initial file size
//	fileInfo, err := file.Stat()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fileSize := fileInfo.Size()
//	reader := bufio.NewReader(file) //buff 4KB
//	for {
//		// Check if the file size has changed
//		newFileInfo, err := file.Stat()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		newFileSize := newFileInfo.Size()
//
//		if newFileSize > fileSize {
//			// New lines appended, read and print them , cp is current position of file
//			cp, err := file.Seek(fileSize, 0)
//			fmt.Printf("the current position is: %v \n", cp)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			scanner := bufio.NewScanner(reader) //gives iterable scanner
//			for scanner.Scan() {
//				fmt.Println(scanner.Text())
//			}
//			/* https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/bufio/scan.go;l=29
//			this shows that if scanner gets EOF a <nil> is set to scanner.Err()
//			*/
//			if scanner.Err() != nil {
//				log.Fatal(scanner.Err())
//			}
//
//			// Update the file size
//			fileSize = newFileSize
//		}
//		time.Sleep(1000 * time.Millisecond)
//	}
//}
//
//func main() {
//	if len(os.Args) != 2 {
//		log.Fatal("gtail <path to file>")
//	}
//
//	filename := os.Args[1]
//
//	gtail(filename)
//}
