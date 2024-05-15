package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Nama Environment bisa diambil dari environment variable
	environmentName := os.Getenv("Environment_Name")

	fmt.Fprintf(w, "<h1>Nama Aplikasi: Test Devsecops Batch-11</h1>")
	fmt.Fprintf(w, "<p>Nama Environment: %s</p>", environmentName)
	fmt.Fprintf(w, "<p>Version: 1.0</p>")
	fmt.Fprintf(w, "<p>Status: OK</p>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}