package main

import (
	"example.com/notes/utils"
)

func main() {
	n := utils.EnterNotesData()
	utils.InsertIntoJson(n)
}
