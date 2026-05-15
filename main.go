package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/internal/handler"
	"todo-api/internal/middleware"
	"todo-api/internal/store"
)

func main() {
	taskStore := store.NewTaskStore()
	taskHandler := handler.NewTaskHandler(taskStore)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("GET /tasks", h.GetAll)
	mux.HandleFunc("GET /tasks", h.GetByID)
	mux.HandleFunc("PUT /tasks", h.Update)
	mux.HandleFunc("DELETE /tasks", h.Delete)

	server := middleware.CORS(
		middleware.Recovery(
			middleware.Logging(mux),
		),
	)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
