package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// handlerChirpsCreate handles POST /api/chirps
func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string `json:"body"`
		UserID string `json:"user_id"`
	}
	type chirpResponse struct {
		ID        string    `json:"id"`
		Body      string    `json:"body"`
		UserID    string    `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if params.Body == "" || params.UserID == "" {
		respondWithError(w, http.StatusBadRequest, "Missing body or user_id", nil)
		return
	}

	id := uuid.New().String()
	createdAt := time.Now().UTC()

	resp := chirpResponse{
		ID:        id,
		Body:      params.Body,
		UserID:    params.UserID,
		CreatedAt: createdAt,
	}

	respondWithJSON(w, http.StatusCreated, resp)
}
