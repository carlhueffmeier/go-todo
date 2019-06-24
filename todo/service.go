package todo

import "fmt"

// Service creates and updates todo items
type Service struct {
	todos []Todo
}

// CreateTodoRequest contains data to create a new todo item
type CreateTodoRequest struct {
	Title string
}

// GetTodosResponse is an array of Todos
type GetTodosResponse []Todo

// NewService creates a new service instance
func NewService() (s Service) {
	return
}

// CreateTodo creates a new todo item
func (s *Service) CreateTodo(r CreateTodoRequest) error {
	t := NewTodo(r.Title)
	s.todos = append(s.todos, t)
	fmt.Printf("New Todo: %q\n", t)
	fmt.Printf("All todos: %v\n", s.todos)
	return nil
}

// GetTodos returns all todos
func (s *Service) GetTodos() (r GetTodosResponse) {
	return s.todos
}
