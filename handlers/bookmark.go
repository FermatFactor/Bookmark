package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/FermatFactor/Bookmark/models"
	"github.com/google/uuid"
)

var (
	bookmarks []models.Bookmark
	mutex     sync.Mutex
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	var b models.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	b.ID = uuid.New().String()

	mutex.Lock()
	bookmarks = append(bookmarks, b)
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func GetAllBookmarks(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmarks)
}

func GetRandomThought(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")

	if len(bookmarks) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"thought": "No thoughts yet!"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := bookmarks[rand.Intn(len(bookmarks))]

	json.NewEncoder(w).Encode(map[string]string{"thought": random.Thought})
}
