package handlers

import (
	"net/http"
	"strconv"

	"github.com/ErayCep/ryg/internal/model"
	"github.com/gorilla/mux"
)

func (h *Handlers) GetReviewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.l.Printf("[INFO] Review Get Handler")

	var reviews model.Reviews

	reviews, err := h.Storage.GetReviews()
	if err != nil {
		http.Error(w, "[ERROR] problem querying reviews from database", http.StatusInternalServerError)
		h.l.Fatal(err)
	}

	reviews.ToJSON(w)
}

func (h *Handlers) GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("[INFO] GET review handler")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.l.Printf("[ERROR] Failed to retrive id from request")
		http.Error(w, "Unable to get review from URL", http.StatusBadRequest)
	}
	var review *model.Review
	review, err = h.Storage.GetReview(id)
	if err != nil {
		h.l.Printf("[ERROR] Failed to get review from database")
		http.Error(w, "Failed to get review from database", http.StatusInternalServerError)
	}

	review.ToJSON(w)
}

func (h *Handlers) PostReviewHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("[INFO] POST review handler")

	review := model.Review{}
	err := review.FromJSON(r.Body)
	if err != nil {
		h.l.Printf("[ERROR] FromJSON failed")
		http.Error(w, "Unable to decode json", http.StatusInternalServerError)
	}

	h.Storage.AddReview(&review)
}

func (h *Handlers) PutReviewHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("[INFO] PUT review handler")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.l.Printf("[ERROR] Failed to get review ID from request")
		http.Error(w, "Unable to retrive ID", http.StatusBadRequest)
	}

	review := model.Review{}
	err = review.FromJSON(r.Body)
	if err != nil {
		h.l.Printf("[ERROR] FromJSON failed")
		http.Error(w, "Unable to decode json", http.StatusInternalServerError)
	}

	h.Storage.UpdateReview(&review, id)
}

func (h *Handlers) DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("[INFO] DELETE review handler")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		h.l.Printf("[ERROR] Failed to get review ID from request body")
		http.Error(w, "Unable to decode json", http.StatusInternalServerError)
	}

	h.Storage.DeleteReview(id)
}
