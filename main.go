package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/FermatFactor/Bookmark/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, reading from system environment")
	}

	// Initialize Supabase client
	InitializeSupabase()

	// Create router
	r := mux.NewRouter()

	// API Routes - prefixed with /api
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/bookmark", handlers.AddBookmark).Methods("POST")
	api.HandleFunc("/bookmarks", handlers.GetAllBookmarks).Methods("GET")
	api.HandleFunc("/thought/random", handlers.GetRandomThought).Methods("GET")

	// Serve static frontend files from ./frontend/
	staticDir := "./frontend"
	absStaticPath, err := filepath.Abs(staticDir)
	if err != nil {
		log.Fatalf("Failed to find static directory: %v", err)
	}
	fs := http.FileServer(http.Dir(absStaticPath))
	r.PathPrefix("/").Handler(fs)

	// Start server
	fmt.Println("🚀 Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
