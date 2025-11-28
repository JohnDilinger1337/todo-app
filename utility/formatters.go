package utility

import "main/models"

func Formatter(todoList *[]models.Todo) models.FormattedTodoList {
	var formattedTodoList models.FormattedTodoList
	for i := range *todoList {
		switch (*todoList)[i].Category {
		case models.Category1:
			formattedTodoList.Category1 = append(formattedTodoList.Category1, (*todoList)[i])
		case models.Category2:
			formattedTodoList.Category2 = append(formattedTodoList.Category2, (*todoList)[i])
		case models.Category3:
			formattedTodoList.Category3 = append(formattedTodoList.Category3, (*todoList)[i])
		}
	}
	return formattedTodoList
}