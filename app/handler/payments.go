package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/achanda/testrest/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllPayments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	payments := []model.Payment{}
	db.Find(&payments)
	respondJSON(w, http.StatusOK, payments)
}

func CreatePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	payment := model.Payment{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, payment)
}

func GetPayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	s, _ := strconv.ParseUint(id, 10, 32)
	payment := getPaymentOr404(db, uint(s), w, r)
	if payment == nil {
		return
	}
	respondJSON(w, http.StatusOK, payment)
}

func UpdatePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	s, _ := strconv.ParseUint(id, 10, 32)
	payment := getPaymentOr404(db, uint(s), w, r)
	if payment == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, payment)
}

func DeletePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	s, _ := strconv.ParseUint(id, 10, 32)
	payment := getPaymentOr404(db, uint(s), w, r)
	if payment == nil {
		return
	}
	if err := db.Delete(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPaymentOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Payment {
	payment := model.Payment{}
	if err := db.Where("ID = ?", id).First(&payment).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &payment
}
