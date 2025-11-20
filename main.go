package main

import (
	"fmt"
	"path/filepath"

	// "github.com/xuri/excelize/v2"
	"os"
)

type Student struct {
	name string
	age int
	email string
	id int
}

type Subject struct {
	teacher string
	title string
	description string
	id int
}

type Grade struct {
	student_id int
	subject_id int
	mark float32
}



func main() {
	entries, err  := os.ReadDir(".")
	if err != nil {
		fmt.Println("Something bad is happening with your computer!")
	}
	var xlsxFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".xlsx" {
			xlsxFiles = append(xlsxFiles, entry.Name())
		}
	}
	fmt.Println("----------Welcome to Students Manager-----------")
	fmt.Println("First you'll choose the file of the list of students: \n")
	for i, file := range xlsxFiles {
		fmt.Printf("%d. %s\n", i+1, file)
	}
	var filenumber int
	fmt.Print("\nPlease enter the number of the right file: ")
	fmt.Scanf("%d", &filenumber)
}