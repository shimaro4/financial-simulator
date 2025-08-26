package main

import (
    "fmt"
	"github.com/shimaro4/financial-simulator/family"
)

func main() {
    family := family.Household{
		Name: "Johnson",
		NumPeople: 4,
		State: "California",
	}

	family.AddIncome("Mike", "Professor", 5500, .05)
	family.AddIncome("Joe", "DEA Agent", 10000, .05)

	fmt.Printf("Family: %v\n", family)
}