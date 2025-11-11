package main

import "fmt"

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
	var students []Student = make([]Student, 0, 20)
	var subjects []Subject = []Subject{
		{title: "Mathematics", id: 1, description: "Advanced concepts of Algebra and Calculus"},
		{teacher: "Peter Machman", title: "Statistics", id: 2, description: "Methods of counting and population analysis"},
	}
	
	students = append(students, Student{name: "Paul", age: 22, id: 10, email: "paulzaimer@gmail.com"})
	var grades []Grade = make([]Grade, 10)
	grades = append(grades, Grade{subject_id: 1, student_id: 10, mark: 14})
	
	for _, grade := range grades {

		var found = 0
		var sub Subject
		for _, sub = range subjects {
			if sub.id == grade.subject_id {
				found ++
				break
			}
		}

		var std Student
		for _, std = range students {
			if std.id == grade.student_id {
				found ++
				break
			}
		}

		if (found>1) {
			fmt.Printf("%s took %.2f in %s\n", std.name, grade.mark, sub.title)
		}
	}
}