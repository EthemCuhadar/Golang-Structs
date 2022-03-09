package main

import (
	"fmt"
)

// Create Point struct
type Point struct{
	x float64
	y float64
}

// Create Employee struct
type Employee struct{
	Name 		string
	Age 		int
	Position 	string
	IsManager	bool
	Salary		float64
}

func main(){
	
	// Define a point whose x and y fields are 0.0
	pointA := Point{0.0, 0.0}
	
	// Print the pointA.
	fmt.Println(pointA)
	fmt.Println(pointA.x)
	fmt.Println(pointA.y)
	
	// Define another point with unfilled fields.
	pointB := Point{}
	
	// Print the pointB.
	fmt.Println(pointB)
	
	// Define an employee
	emp001 := Employee{
		Name: "Elliot Anderson",
		Age: 29,
		Position: "Engineer",
		IsManager: false,
		Salary: 4000,
	}
	
	// Print emp001 and fields.
	fmt.Println(emp001)
	
	x := &emp001
	x.Salary = 5000
	fmt.Println(emp001)
	
	// ANONYMOUS STRUCTS
	
	// Define an anonymous struct.
	manager001 := struct{
		Name		string
		Age			int
		Position	string
		IsManager	bool
		Salary		float64
	}{
		Name: "Tyrell Wellick",
		Age: 32,
		Position: "CTO",
		IsManager: true,
		Salary: 8000,
	}
	
	// Print the anonymous struct.
	fmt.Println(manager001)
	
	// EQUALITIES
	
	// Define two employees
	emp002 := Employee{
		Name: "Elliot Anderson",
		Age: 29,
		Position: "Engineer",
		IsManager: false,
		Salary: 5000,
	}
	
	emp003 := Employee{
		Name: "Elliot Robertson",  // Name is different
		Age: 29,
		Position: "Engineer",
		IsManager: false,
		Salary: 4000,
	}
	
	// Print the equalities
	fmt.Println(emp001 == emp002)
	fmt.Println(emp002 == emp003)
}
