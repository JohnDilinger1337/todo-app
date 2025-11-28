package main

import (
	"fmt"
	"main/models"
	"main/utility"
)

func main() {
	todos := []models.Todo{
		{ID: 1, Category: "Category1", Title: "Learn Go", Done: false},
		{ID: 2, Category: "Category2", Title: "Build a web app", Done: false},
		{ID: 3, Category: "Category1", Title: "Deploy to cloud", Done: false},
		{ID: 4, Category: "Category3", Title: "Fix bugs", Done: false},
	}

	formattedTodoList := utility.Formatter(&todos)

	fmt.Println("############# Welcome to my to-do list #############")

	fmt.Printf("%s {\n", models.Category1)
	for i := range formattedTodoList.Category1 {
		fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category1[i].Title, utility.DoneMessage(formattedTodoList.Category1[i].Done))
	}
	fmt.Printf("}\n%s {\n", models.Category2)
	for i := range formattedTodoList.Category2 {
		fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category2[i].Title, utility.DoneMessage(formattedTodoList.Category2[i].Done))
	}
	fmt.Printf("}\n%s {\n", models.Category3)
	for i := range formattedTodoList.Category3 {
		fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category3[i].Title, utility.DoneMessage(formattedTodoList.Category3[i].Done))
	}
	fmt.Println("}")
}
