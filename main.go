package main

import (
	"fmt"
	"strconv"
)

type Category string

const (
	Category1 	Category = "Category1"
	Category2 	Category = "Category2"
	Category3 	Category = "Category3"
)

type FormattedTodoList struct {
    Category1 []Todo
    Category2 []Todo
    Category3 []Todo
}

type Todo struct {
	ID    int
	Category Category
	Title string
	Done  bool
}

type TodoInput struct {
	Title string
	Category Category
}

func main() {
	todos := []Todo{
		{ID: 1, Category: "Category1", Title: "Learn Go", Done: false},
		{ID: 2, Category: "Category2", Title: "Build a web app", Done: false},
		{ID: 3, Category: "Category1", Title: "Deploy to cloud", Done: false},
		{ID: 4, Category: "Category3", Title: "Fix bugs", Done: false},
	}

	var formattedTodoList FormattedTodoList

	// const greeting = "Hello, World!"

	fmt.Println("############# Welcome to my to-do list #############")
	for i := range todos {
		switch todos[i].Category {
		case Category1:
			formattedTodoList.Category1 = append(formattedTodoList.Category1, todos[i])
		case Category2:
			formattedTodoList.Category2 = append(formattedTodoList.Category2, todos[i])
		case Category3: 
			formattedTodoList.Category3 = append(formattedTodoList.Category3, todos[i])
		} 
	}

	fmt.Printf("%s {\n", Category1)
	for i := range formattedTodoList.Category1 {
		fmt.Printf("%v) %v (%v)\n", i +1, formattedTodoList.Category1[i].Title, doneMessage(formattedTodoList.Category1[i].Done))
	}
	fmt.Printf("}\n%s {\n", Category2)
	for i := range formattedTodoList.Category2 {
		fmt.Printf("%v) %v (%v)\n", i +1, formattedTodoList.Category2[i].Title, doneMessage(formattedTodoList.Category2[i].Done))
	}	
	fmt.Printf("}\n%s {\n", Category3)
	for i := range formattedTodoList.Category3 {
		fmt.Printf("%v) %v (%v)\n", i +1, formattedTodoList.Category3[i].Title, doneMessage(formattedTodoList.Category3[i].Done))
	}
	fmt.Println("}")
}

func (c Category) isValid() bool {
	switch c {
	case Category1, Category2, Category3:
		return true
	}
	return false
}

func doneMessage(done bool) string {
	message := map[bool]string{
		true: "This is already done!",
		false: "Not done yet!",
	}
	return message[done]
	
}

 // ???? Test to change %s 
func addNewItem(item TodoInput, todos *[]Todo) error {
	if item.Title == ""{return fmt.Errorf("invalid title: %s", item.Title)} else if !item.Category.isValid() {return fmt.Errorf("invalid category: %s", item.Category)}
	*todos = append(*todos, Todo{ID: 5, Title: item.Title, Done: false})
	return nil
}

func deleteItem(ID int, todos *[]Todo) error {
	if todos == nil || ID <= 0 || ID > len(*todos) {
		return fmt.Errorf("something went wrong while deleteing this item with ID: (%s)", strconv.Itoa(ID))
	}
	*todos = append((*todos)[:ID-1], (*todos)[ID:]...)

	return nil
}
