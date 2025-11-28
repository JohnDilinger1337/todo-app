package services

import (
	"fmt"
	"main/models"
	"strconv"
)

func AddNewItem(input models.TodoInput, todo *[]models.Todo) error {
	if input.Title == "" {
		return fmt.Errorf("invalid title: %s", input.Title)
	} else if !input.Category.IsValid() {
		return fmt.Errorf("invalid category: %s", input.Category)
	}

	item := models.Todo{
		ID:       len(*todo) + 1,
		Title:    input.Title,
		Category: input.Category,
		Done:     false,
	}

	*todo = append(*todo, item)

	return nil
}

func DeleteNewItem(input int, todo *[]models.Todo) error {
	if *todo == nil || input <= 0 || input > len(*todo) {
		return fmt.Errorf("something went wrong while deleteing this item with ID: (%s)", strconv.Itoa(input))
	}

	*todo = append((*todo)[:input-1], (*todo)[input:]...)

	return nil
}
