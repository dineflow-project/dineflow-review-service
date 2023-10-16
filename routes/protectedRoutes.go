package routes

import (
	"net/http"

	"dineflow-review-services/controllers"

	"github.com/gorilla/mux"
)

func ProtectedRoute(r *mux.Router) {
	r.HandleFunc("/canteens", controllers.GetAllReviews).Methods("GET")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.GetReviewByID).Methods("GET")
	r.HandleFunc("/canteens", controllers.CreateReview).Methods("POST")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.UpdateReviewByID).Methods("PUT", "PATCH")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.DeleteReviewByID).Methods("DELETE")

	http.Handle("/", r)
}
