package main

import (
	"GoProgrammingRemote/employee"
	"fmt"
)

func main() {
	e := employee.Employee{
		FirstName:   "Maynor",
		LastName:    "Teleguario",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
	e.PrintName()
	e.UpdateName("Gustavo", "Queche")
	fmt.Println("Name was Update to")
	e.PrintName()
}
