package main

import (
	"log"
	"net/http"
	"project/pkg"
)

func main() {
	service := pkg.NewService()
	handler := pkg.NewHandler(service)

	http.HandleFunc("/tasks", handler.CreateTask)
	http.HandleFunc("/tasks/", handler.GetTask)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
