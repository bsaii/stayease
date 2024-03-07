package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bsaii/stayease/internal/model"
	"github.com/bsaii/stayease/internal/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func AddBill(w http.ResponseWriter, r *http.Request) {
	bill := &model.Bill{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bill); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Create(&bill).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating record for bill: %s", err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, bill)
}

func GetBill(w http.ResponseWriter, r *http.Request) {
	bill := &model.Bill{}

	billId := chi.URLParam(r, "billID")
	billIdInt, err := strconv.Atoi(billId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&bill, billIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error finding bill with id %s: %s", billId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, bill)
}

func AllBills(w http.ResponseWriter, r *http.Request) {
	bills := []model.Bill{}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Find(&bills).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error finding all bills: %s", err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, bills)
}

func UpdateBill(w http.ResponseWriter, r *http.Request) {
	bill := &model.Bill{}
	bill_update := &model.Bill{}

	billId := chi.URLParam(r, "billID")
	billIdInt, err := strconv.Atoi(billId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bill_update); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&bill, billIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error finding bill with id %s: %s", billId, err.Error()))
		return
	}

	if bill_update.Amount > 0 && bill_update.Amount != bill.Amount {
		bill.Amount = bill_update.Amount
	}

	if bill_update.CustomerName != "" {
		bill.CustomerName = bill_update.CustomerName
	}

	if bill_update.Description != "" {
		bill.Description = bill_update.Description
	}

	if bill_update.Paid != bill.Paid {
		bill.Paid = bill_update.Paid
	}

	if err := db.Save(&bill).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating bill with id %s: %s", billId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, bill)
}

func DelBill(w http.ResponseWriter, r *http.Request) {
	bill := &model.Bill{}

	billId := chi.URLParam(r, "billID")
	billIdInt, err := strconv.Atoi(billId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Delete(&bill, billIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting room with id %s: %s", billId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, fmt.Sprintf("Successfully deleted bill with Id: %s", billId))
}
