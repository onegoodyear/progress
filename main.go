package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
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
	grade_id int
	student_id int
	subject_id int
	mark float32
}



func main() {
	// var grades []Grade = []Grade{}
	var students []Student = []Student{}
	// var subjects []Subject = []Subject{}
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
	fmt.Println("First you'll choose the file of the list of students:")
	for i, file := range xlsxFiles {
		fmt.Printf("\n%d. %s\n", i+1, file)
	}
	var filenumber int
	fmt.Print("\nPlease enter the number of the right file: ")
	fmt.Scanf("%d", &filenumber)

	// Reading the List of students

	studentsFile, err := excelize.OpenFile(xlsxFiles[filenumber-1])
	if err != nil {
		fmt.Println("Error opeing the students file: ", err)
	}
	rows, err := studentsFile.GetRows("Sheet1")
	if err != nil {
		fmt.Println("There was an error getting the rows of the students file")
	}
	for i:=1; i < len(rows); i++ {
		row := rows[i]
		ageInt, _ := strconv.Atoi(row[1])
		idInt, _ := strconv.Atoi(row[3])
		students = append(students, Student{
			name: row[0],
			age: ageInt,
			email: row[2],
			id: idInt,
		})
	}
	fmt.Println(students)
}