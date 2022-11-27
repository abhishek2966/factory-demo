package main

import (
	"fmt"

	"github.com/abhishek2966/factory-demo/pkg/resource"
)

func main() {
	empGeneratorBackend := resource.NewEmployeeGenerator(resource.BackendTeam)
	b1 := empGeneratorBackend("John", "john@email.com")
	b2 := empGeneratorBackend("Sam", "sam@email.com")
	fmt.Printf("%+v\n", b1) // {Name:John Email:john@email.com Team:backend}
	fmt.Printf("%+v\n", b2) // {Name:Sam Email:sam@email.com Team:backend}

	empGeneratorUI := resource.NewEmployeeGenerator(resource.FrontendTeam)
	f1 := empGeneratorUI("Jayesh", "Jayesh@email.com")
	f2 := empGeneratorUI("Kyle", "kyle@email.com")
	fmt.Printf("%+v\n", f1) // {Name:Jayesh Email:Jayesh@email.com Team:frontend}
	fmt.Printf("%+v\n", f2) // {Name:Kyle Email:kyle@email.com Team:frontend}
}
