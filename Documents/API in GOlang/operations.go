package main

import (
	"fmt"
	"math"
	"os"
)

func addStudent() {
	if dataBase == nil {
		dataBase = make(map[string]Student)
	}
	fmt.Println("Enter the first name of the student:")
	var fname string
	fmt.Scanln(&fname)

	fmt.Println("Enter the last name of the student:")
	var lname string
	fmt.Scanln(&lname)

	fmt.Println("Enter the PRN of the student:")
	var prn string
	fmt.Scanln(&prn)

	fmt.Println("Enter the marks of the student:")
	var marks int
	fmt.Scanln(&marks)

	student := Student{fname, lname, prn, marks}

	fmt.Println("Student added successfully!")
	fmt.Println()
	// fmt.Println(student)
	dataBase[prn] = student
}

func viewStudents() {
	for prn, student := range dataBase {
		fmt.Println("PRN:", prn)
		fmt.Println("Name:", student.fname+" "+student.lname)
		fmt.Println("Marks:", student.marks)
		fmt.Println()
	}
	fmt.Println()
}

func updateStudent() {
	fmt.Println("Enter the PRN of the student to be updated:")
	var prn string
	fmt.Scanln(&prn)
	_, exists := dataBase[prn]
	if exists {
		fmt.Println("Enter the first name of the student:")
		var fname string
		fmt.Scanln(&fname)

		fmt.Println("Enter the last name of the student:")
		var lname string
		fmt.Scanln(&lname)

		fmt.Println("Enter the marks of the student:")
		var marks int
		fmt.Scanln(&marks)

		student := Student{fname, lname, prn, marks}
		dataBase[prn] = student
		fmt.Println("Student updated successfully!")
		fmt.Println()
	} else {
		fmt.Println("Student not found!")
		fmt.Println()
	}

}

func deleteStudent() {
	fmt.Println("Enter the PRN of the student to be deleted:")
	var prn string
	fmt.Scanln(&prn)
	_, exists := dataBase[prn]
	if exists {
		delete(dataBase, prn)
		fmt.Println("Student deleted successfully!")
		fmt.Println()
	} else {
		fmt.Println("Student not found!")
		fmt.Println()
	}
}

func exportToTxt() {
	// start := time.Now()
	fmt.Println("Enter the name of the file:")
	var filename string
	fmt.Scanln(&filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for prn, student := range dataBase {
		fmt.Fprintf(f, "PRN: %s\n", prn)
		fmt.Fprintf(f, "Name: %s\n", student.fname+" "+student.lname)
		fmt.Fprintf(f, "Marks: %d\n\n", student.marks)
	}

	// fmt.Println("Time taken to export:", time.Since(start))

	fmt.Println("Data exported successfully!")
	fmt.Println()
}

func getSum() int {
	sum := 0
	for _, student := range dataBase {
		sum += student.marks
	}
	return sum
}

func getStats() (sum int, avg float64, highest int, lowest int) {
	sum = getSum()
	avg = float64(sum) / float64(len(dataBase))
	highest = 0
	lowest = math.MaxInt64
	for _, student := range dataBase {
		if student.marks > highest {
			highest = student.marks
		}
		if student.marks < lowest {
			lowest = student.marks
		}
	}

	return
}