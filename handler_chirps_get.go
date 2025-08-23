package main

import (
	"net/http"
	"time"
)

// Chirp is the API representation of a chirp
type Chirp struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID.String(),
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID.String(),
			Body:      dbChirp.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
