package todo

import "fmt"

// Todo represents a todo item
type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// NewTodo creates a new todo
func NewTodo(title string) (t Todo) {
	t.Title = title
	return
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo %q (Done=%v)", t.Title, t.Done)
}
