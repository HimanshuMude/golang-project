package main

import (
	"fmt"
)

type Student struct {
	fname string
	lname string
	prn   string
	marks int
}

var dataBase map[string]Student



func main() {

	fmt.Println("Welcome to the Student DB!")
	fmt.Println()
	for {
		
		loadChoices()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			viewStudents()
		case 3:
			updateStudent()
		case 4:
			deleteStudent()
		case 5:
			exportToTxt()
		case 6:
			sum,avg,highest,lowest:=getStats()
			fmt.Println("Sum:",sum)
			fmt.Println("Average:",avg)
			fmt.Println("Highest:",highest)
			fmt.Println("Lowest:",lowest)
			fmt.Println()
		case 7:
			fmt.Println("Exiting...")
			fmt.Println()
			return
		default:
			fmt.Println("Invalid choice!")
			fmt.Println()
		}
	}

}
