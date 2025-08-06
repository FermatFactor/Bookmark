package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FermatFactor/Bookmark/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Bookmark API!")
	}).Methods("GET")

	r.HandleFunc("/bookmark", handlers.AddBookmark).Methods("POST")
	r.HandleFunc("/bookmarks", handlers.GetAllBookmarks).Methods("GET")
	r.HandleFunc("/thought/random", handlers.GetRandomThought).Methods("GET")

	fmt.Println("🚀 Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
