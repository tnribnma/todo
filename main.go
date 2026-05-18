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
	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks", taskHandler.GetAll)
	mux.HandleFunc("GET /tasks/", taskHandler.GetByID)
	mux.HandleFunc("PUT /tasks/", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/", taskHandler.Delete)

	server := middleware.CORS(
		middleware.Recovery(
			middleware.Logging(mux),
		),
	)

	fmt.Println("Todo API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
