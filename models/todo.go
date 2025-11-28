package models

type Category string

const (
	Category1 Category = "Category1"
	Category2 Category = "Category2"
	Category3 Category = "Category3"
)

type FormattedTodoList struct {
	Category1 []Todo
	Category2 []Todo
	Category3 []Todo
}

type Todo struct {
	ID       int
	Category Category
	Title    string
	Done     bool
}

type TodoInput struct {
	Title    string
	Category Category
}

func (c Category) IsValid() bool {
	switch c {
	case Category1, Category2, Category3:
		return true
	}
	return false
}