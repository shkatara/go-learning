package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Notes struct {
	Title   string
	Content string
}

const fileName = "test.txt"

func EnterNotesData() Notes {
	var title string
	var content string
	fmt.Print("Enter Notes Title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter Notes Content: ")
	fmt.Scanln(&content)
	return Notes{
		Title:   title,
		Content: content,
	}
}

func InsertIntoJson(n Notes) error {
	data, _ := json.Marshal(n)
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	n := EnterNotesData()
	InsertIntoJson(n)
}
