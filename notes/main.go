package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	//err := os.WriteFile(fileName, data, 0644)  //Always truncates the file. Seems to be overwriting the existing data of file.
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //Using os.O_APPEND syscall to stop overwriting and open the file in append mode
	if err != nil {
		log.Fatal(err)
		return err
	}
	if _, err := f.Write(data); err != nil { //Shorthand like python. I like python better here.
		log.Fatal(err)
		return err
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	n := EnterNotesData()
	InsertIntoJson(n)
}
