package routes

import (
	"net/http"

	"dineflow-review-service/controllers"

	"github.com/gorilla/mux"
)

func ProtectedRoute(r *mux.Router) {
	r.HandleFunc("/reviews", controllers.GetAllReviews).Methods("GET")
	r.HandleFunc("/reviews/{_id}", controllers.GetReviewByID).Methods("GET")
	r.HandleFunc("/reviews", controllers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{_id}", controllers.UpdateReviewByID).Methods("PUT", "PATCH")
	r.HandleFunc("/reviews/{_id}", controllers.DeleteReviewByID).Methods("DELETE")

	http.Handle("/", r)
}
