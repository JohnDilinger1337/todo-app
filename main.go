package main

import (
	"fmt"
	"main/models"
	"main/services"
	"main/utils"
)

const (
	ViewItemsOption  = 1
	AddItemOption    = 2
	DeleteItemOption = 3
	ExitOption       = 4
)

func main() {
	todoSvc := services.NewTodoService()

	for {
		utils.ClearScreen()

		fmt.Println("############# Welcome to my to-do list #############")
		fmt.Println("1) View all items")
		fmt.Println("2) Add new item")
		fmt.Println("3) Delete an item")
		fmt.Println("4) Exit")

		option := utils.PromptInt("Please pick an option of what you need: ")

		switch option {
		case ViewItemsOption:
			todo := todoSvc.GetAll()
			formattedTodoList := utils.Formatter(&todo)

			fmt.Printf("%s {\n", models.Category1)
			for i := range formattedTodoList.Category1 {
				fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category1[i].Title, utils.DoneMessage(formattedTodoList.Category1[i].Done))
			}
			fmt.Printf("}\n%s {\n", models.Category2)
			for i := range formattedTodoList.Category2 {
				fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category2[i].Title, utils.DoneMessage(formattedTodoList.Category2[i].Done))
			}
			fmt.Printf("}\n%s {\n", models.Category3)
			for i := range formattedTodoList.Category3 {
				fmt.Printf("%v) %v (%v)\n", i+1, formattedTodoList.Category3[i].Title, utils.DoneMessage(formattedTodoList.Category3[i].Done))
			}
			fmt.Println("}")

		case AddItemOption:
			title := utils.PromptText("Enter the title of the to-do item: ")
			categoryInput := utils.PromptText("Enter the category: " + string(models.Category1) + ", " + string(models.Category2) + ", " + string(models.Category3) + ": ")
			category := models.Category(categoryInput)
			todoInput := models.TodoInput{
				Title:    title,
				Category: category,
			}
			err := todoSvc.Add(todoInput)
			if err != nil {
				fmt.Printf("Error adding to-do item: %v\n", err)
			}

		case DeleteItemOption:
			id := utils.PromptInt("Enter to-do item's ID: ")

			err := todoSvc.Delete(id)

			if err != nil {
				fmt.Printf("Error deleteing to-do item: %v\n", err)
			}

		case ExitOption:
			fmt.Println("############# Thank you for using my to-do list #############")
			return
		}
		utils.PromptText("Press Enter to continue...")
	}

}
