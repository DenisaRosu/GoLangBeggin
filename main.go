package main

import (
	structure "AtossGoLang/entity"
	"fmt"
)

func main() {
	//intern := structure.Employee{Name: "Ana"}

	//fmt.Println(intern)
	companies := structure.InitCompanies()
	fmt.Println(CountofEmployees(companies, "Atoss", "Timisoara"))
}
func CountofEmployees(companies []structure.Company, name string, city string) int {
	count := 0
	for _, comp := range companies {
		if comp.Name == name{
		
		for _, empl := range comp.Employees{
		  if empl.GetCity() == city{
			count++
		}
	}
  }
}
	return count
}
