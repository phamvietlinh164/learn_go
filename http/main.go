package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	fmt.Println("Server is running...!")

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/api/music", MusicApi)
	http.HandleFunc("/api/students", StudentsApi)

	log.Fatal(http.ListenAndServe(":8200", nil))
}

func StudentsApi(w http.ResponseWriter, r *http.Request) {
	var data = []Student{
		{1, "A", 18},
		{2, "B", 18},
		{3, "C", 18},
	}
	json.NewEncoder(w).Encode(data)
}

func MusicApi(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"name":   "hello",
		"singer": "a",
	}
	json.NewEncoder(w).Encode(data)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is home page</h1>")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is about page</h1>")
}
