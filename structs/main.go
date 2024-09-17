package main

import "structs/employee"

func main() {
	e := employee.Employee {
		FirstName: "Sam",
		LastName: "Suffit",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
