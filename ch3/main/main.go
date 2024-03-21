package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
}

func SaveBooks(fileName string, books []Book) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(string(jsonData) + "\n")
	}
	return writer.Flush()
}

func LoadBooks(fileName string) ([]Book, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var books []Book
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var book Book
		err = json.Unmarshal(scanner.Bytes(), &book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, scanner.Err()
}

func main() {
	fileName := "books.json"
	books := []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 380},
		{"Go in Action", "William Kennedy", 300},
	}

	err := SaveBooks(fileName, books)
	if err != nil {
		fmt.Println("error saving ...")
		return
	}

	loadedBooks, err := LoadBooks(fileName)
	if err != nil {
		fmt.Println("error loading ...")
		return
	}

	fmt.Println(loadedBooks)
}
