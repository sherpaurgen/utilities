package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

const (
	dbPath         = "checkpoint.db"
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS log_position (
			id INTEGER PRIMARY KEY,
			filename TEXT,
			position INTEGER
		);
	`
)

func createDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getLastPosition(db *sql.DB, filename string) (int64, error) {
	var position int64
	query := "SELECT position FROM log_position WHERE filename = ? LIMIT 1;"
	err := db.QueryRow(query, filename).Scan(&position)
	return position, err
}

func updateLastPosition(db *sql.DB, filename string, position int64) error {
	query := "REPLACE INTO log_position (id, filename, position) VALUES (1, ?, ?);"
	_, err := db.Exec(query, filename, position)
	return err
}

func gtail(filename string) {
	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	lastPosition, err := getLastPosition(db, filename)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file) //buff 4KB
	_, err = file.Seek(lastPosition, 0)

	for {
		// Check if the file size has changed
		newFileInfo, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		newFileSize := newFileInfo.Size()

		if newFileSize > fileSize {
			// New lines appended, read and print them , cp is current position of file
			cp, err := file.Seek(fileSize, 0)
			fmt.Printf("the current position is: %v \n", cp)
			if err != nil {
				log.Fatal(err)
			}

			scanner := bufio.NewScanner(reader) //gives iterable scanner
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
			err = updateLastPosition(db, filename, cp) //update last read position info in sqlite db
			if err != nil {
				log.Fatal("updateLastPosition:", err)
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("gtail <path to file>")
	}

	filename := os.Args[1]

	gtail(filename)
}
