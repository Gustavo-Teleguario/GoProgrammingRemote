package main

import (
	"GoProgrammingRemote/employee"
)

func main() {
	e := employee.Employee{
		FirstName:   "Maynor",
		LastName:    "Teleguario",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
