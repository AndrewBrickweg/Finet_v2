package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"message": "Portfolio API is working!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)	
}