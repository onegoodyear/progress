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
	id int
	student_id int
	subject_id int
	mark float32
}

func displayFiles (files []string) {
	for i, file := range files {
		fmt.Printf("\n %d. %s", i+1, file)
	}
}

func main() {
	var grades []Grade = []Grade{}
	var students []Student = []Student{}
	var subjects []Subject = []Subject{}
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
	displayFiles(xlsxFiles)
	var filenumber int
	for {
		fmt.Print("\nPlease enter the number of the right file: ")
		_, err = fmt.Scanln(&filenumber)
		if err == nil && filenumber > 0 && filenumber <= len(xlsxFiles) {
			break
		}
		fmt.Println("Invalid Input. Try Again!!")
	}


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

	studentsFile.Close()

	// Reading the list of subjects

	fmt.Println("Second you'll choose the file of the list of subjects:")
	displayFiles(xlsxFiles)
	for {
		fmt.Print("\nPlease enter the number of the right file: ")
		_, err = fmt.Scanln(&filenumber)
		if err == nil && filenumber > 0 && filenumber <= len(xlsxFiles) {
			break
		}
		fmt.Println("Invalid Input. Try Again!!")
	}

	subjectsFile, err := excelize.OpenFile(xlsxFiles[filenumber-1])
	if err != nil {
		fmt.Println("There was a problem opening the students file, ", err)
	}

	rows, err = subjectsFile.GetRows("Sheet1")
	if err != nil {
		fmt.Println("There was a problem opening the subjects file, ", err)
	}
	for i := 1; i<len(rows); i++ {
		row := rows[i]
		idInt, _ := strconv.Atoi(row[3])
		subjects = append(subjects, Subject{
			id: idInt,
			title: row[1],
			description: row[2],
			teacher: row[0],
		})
	}

	// Reading the list of the grades

	fmt.Println("Now it's the time to select the grades file: ")
	displayFiles(xlsxFiles)
	for {
		fmt.Println("\nPlease enter the right number: ")
		_, err := fmt.Scanln(&filenumber)
		if err == nil && filenumber > 0 && filenumber <= len(xlsxFiles) {
			break
		}
		fmt.Println("Invalid Input, Try Again.")
	}
	
	gradesFile, err := excelize.OpenFile(xlsxFiles[filenumber-1])
	if err != nil {
		fmt.Println("There was a problem opening the grades file, ", err)
	}
	rows, err = gradesFile.GetRows("Sheet1")
	if err != nil {
		fmt.Println("There were a problem getting the rows of the grades file")
	}
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		idInt, _ := strconv.Atoi(row[0])
		studentId, _ := strconv.Atoi(row[1])
		subjectId, _ := strconv.Atoi(row[2])
		mark64, _ := strconv.ParseFloat(row[3], 32)
		mark := float32(mark64)
		grades = append(grades, Grade {
			id: idInt,
			student_id: studentId,
			subject_id: subjectId,
			mark: mark,
		})
	}

	// Writting the results in output.txt

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("An error occured when creating the output file, ", err)
	}
	outputFile.WriteString("--------The List of Grades (written)-------\n")

	for i, grade := range grades {
		var studentName string
		for _, student := range students {
			if student.id == grade.student_id {
				studentName = student.name
				break
			}
		}
		var subjectTitle string
		for _, subject := range subjects {
			if (subject.id == grade.subject_id) {
				subjectTitle = subject.title
				break
			}
		}
		outputFile.WriteString(fmt.Sprintf("%d. %s took %f on: %s\n", i+1, studentName, grade.mark, subjectTitle))
	}

}