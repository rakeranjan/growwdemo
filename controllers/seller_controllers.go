package controllers

import (
	"encoding/json"
	items "groww/Items"
	"groww/requests"
	"groww/sellers"
	"net/http"
)

func AddSeller(w http.ResponseWriter, r *http.Request) {
	seller := sellers.AddSeller()
	if seller != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seller)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ListItem(w http.ResponseWriter, r *http.Request) {
	request := requests.ListItem{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	seller := sellers.FindSeller(request.SellerID)
	if seller == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := seller.AddItem(request.Name, request.Price)
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
	w.Header().Set("Content-Type", "application/json")
}

func GetAllListing(w http.ResponseWriter, r *http.Request) {
	items := items.GetAllItem(false)
	if items == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(items)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetSellerListings(w http.ResponseWriter, r *http.Request) {
	request := requests.GetSellerListingsRequests{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	seller := sellers.FindSeller(request.SellerID)
	if seller == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(seller.Items)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
