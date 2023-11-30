package controllers

import (
	"encoding/json"
	"groww/buyers"
	"groww/requests"
	"net/http"
)

func AddBuyer(w http.ResponseWriter, r *http.Request) {
	buyer := buyers.AddBuyer()
	if buyer != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(buyer)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func Buy(w http.ResponseWriter, r *http.Request) {
	request := requests.BuyRequests{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	buyer := buyers.FetchBuyerByID(request.BuyerID)
	if buyer == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := buyer.Buy(request.ItemID)
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buyer)
	w.Header().Set("Content-Type", "application/json")
}

func AddToWishlist(w http.ResponseWriter, r *http.Request) {
	request := requests.BuyRequests{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	buyer := buyers.FetchBuyerByID(request.BuyerID)
	if buyer == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := buyer.AddToWishlist(request.ItemID)
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buyer)
	w.Header().Set("Content-Type", "application/json")
}
