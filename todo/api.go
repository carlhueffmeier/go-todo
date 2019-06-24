package todo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// API is responsible for serving requests to our Todo API
type API struct {
	s *Service
}

// NewAPI Creates a new API
func NewAPI(s *Service) (api API) {
	api.s = s
	return
}

func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	if method == "POST" && path == "" {
		request := parseCreateTodoRequest(r.Body)
		if err := api.s.CreateTodo(request); err != nil {
			log.Fatal(err)
		}

	} else if method == "GET" && path == "" {
		todos := api.s.GetTodos()
		jsonRes, err := json.Marshal(todos)
		if err != nil {
			log.Fatalf("Error marshalling response: %s", err)
		}
		if _, err = w.Write(jsonRes); err != nil {
			log.Printf("Error writing response: %s", err)
		}
	}
}

func parseCreateTodoRequest(body io.ReadCloser) (r CreateTodoRequest) {
	dec := json.NewDecoder(body)
	if err := dec.Decode(&r); err != nil {
		body.Close()
		log.Fatal(err)
	}
	body.Close()
	return
}
