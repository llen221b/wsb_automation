package main

import (
	"encoding/json"
	"net/http"
)

type IncomeRequest struct {
	Revenue float64 `json:"revenue"`
	Costs   float64 `json:"costs"`
}

type IncomeResponse struct {
	Profit float64 `json:"profit"`
}

func calculateProfit(w http.ResponseWriter, r *http.Request) {
	var req IncomeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := IncomeResponse{
		Profit: req.Revenue - req.Costs,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/calculate", calculateProfit)
	http.ListenAndServe(":8080", nil)
}
