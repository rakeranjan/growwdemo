package server

import (
	"groww/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/addSeller", Interceptor(controllers.AddSeller)).Methods("POST")
	r.HandleFunc("/addBuyer", Interceptor(controllers.AddBuyer)).Methods("POST")
	r.HandleFunc("/listItem", Interceptor(controllers.ListItem)).Methods("POST")
	r.HandleFunc("/getAllListing", Interceptor(controllers.GetAllListing)).Methods("GET")
	r.HandleFunc("/getSellerListing", Interceptor(controllers.GetSellerListings)).Methods("POST")
	r.HandleFunc("/buy", Interceptor(controllers.Buy)).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func Interceptor(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.RequestURI)
		f(w, r)
	})
}
