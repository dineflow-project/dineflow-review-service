package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dineflow-review-service/models"

	"github.com/gorilla/mux"
)

func GetAllReviews(w http.ResponseWriter, r *http.Request) {
	results, err := models.GetAllReviews()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID := vars["_id"]

	// Query the database to get the  by ID using the new function
	review, err := models.GetReviewByID(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func GetReviewByVendorID(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["_id"]

	// Query the database to get the by vendor ID using the new function
	review, err := models.GetReviewByVendorID(vendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func GetAvgReviewScoreByVendorID(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["_id"]

	avg_score, err := models.GetAvgReviewScoreByVendorID(vendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avg_score)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var newReview models.Review
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newReview); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := models.CreateReview(newReview)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Review created successfully")
}

func DeleteReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID := vars["_id"]

	err := models.DeleteReviewByID(reviewID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Review deleted successfully")
}

func UpdateReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID := vars["_id"]

	var updatedReview models.Review
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedReview); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := models.UpdateReviewByID(reviewID, updatedReview)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Review updated successfully")
}
