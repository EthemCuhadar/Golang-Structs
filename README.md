# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

# Structs

* Structs are another composite type in Go programming language. A **struct** can be defined as collection of fields. These fields belong to the struct that is created.

* Structs are user-defined types which allows gather different types of items into a single item. 

* Let's create a simple struct.

```go
package main

import "fmt"

type Point struct{
    x float64
    y float64
}

func main(){
    pointA := Point{0.0, 0.0}
    fmt.Println(pointA)
    fmt.Println(pointA.x)
    fmt.Println(pointA.y)
}
```

```console
{0 0}
0
0
```

* As you can see, we have created **Point struct** whose fields are **x** and **y**. Moreover, we can access any field by using dot.
* If struct field is not filled, Go programming language sees that field as zero.

```go
	// Define another point with unfilled fields.
	pointB := Point{}
	
	// Print the pointB.
	fmt.Println(pointB)
```

```console
{0 0}
```

-----------------------------------------

## Pointers to structs

* Go programming language allows the users to access any field of a struct by using pointers without dereferencing. That is very practical in usage. 

* On the other hand, address operator, **&**, can be use to specify the address of structs.

```go
// Create Employee struct
type Employee struct{
    Name         string
    Age         int
    Position     string
    IsManager    bool
    Salary        float64
}
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
```

```console
{Elliot Anderson 29 Engineer false 4000}
{Elliot Anderson 29 Engineer false 5000}
```

--------------------------------------

## Anonymous structs

* Sometimes we want to decleare structs without creating a new data type. These structs are called **anonymous structs**.

```go
package main

import "fmt"

func main(){

    // Define an anonymous struct.
    manager001 := struct{
        Name        string
        Age            int
        Position    string
        IsManager    bool
        Salary        float64
    }{
        Name: "Tyrell Wellick",
        Age: 32,
        Position: "CTO",
        IsManager: true,
        Salary: 8000,
    }

    // Print the anonymous struct.
    fmt.Println(manager001)
}
```

```console
{Tyrell Wellick 32 CTO true 8000}
```

----------------------------------

## Struct types and equalities

```go
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
```

```console
true
false
```

* In this example **emp002** has 5000 salary. In the first part **emp001** initially defined with 4000 salary but it changed afterwards. The equality for these two variables are equal at the end.

* However, **emp003** has different name and different salary. Therefore, they are not equal.
