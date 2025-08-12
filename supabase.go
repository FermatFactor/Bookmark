package main

import (
	"log"
	"os"

	supabase "github.com/supabase-community/supabase-go"
)

var Supabase *supabase.Client

// InitializeSupabase sets up the global Supabase client
func InitializeSupabase() {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_ANON_KEY")

	if url == "" || key == "" {
		log.Fatal("❌ Missing SUPABASE_URL or SUPABASE_ANON_KEY in environment variables")
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Supabase client: %v", err)
	}

	Supabase = client
	log.Println("✅ Supabase client initialized successfully")
}
