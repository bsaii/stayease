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

func AddBookingWithRoomId(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}
	booking := &model.Booking{}

	roomId := chi.URLParam(r, "roomID")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&room, roomIdInt).Association("BookedDates").Append(&booking); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Failed to add booking to room with id %s: %s", roomId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, booking)
}

func GetBookingsWithRoomId(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}
	bookings := []model.Booking{}

	roomId := chi.URLParam(r, "roomID")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&room, roomIdInt).Association("BookedDates").Find(&bookings); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Failed to find bookings to room with id %s: %s", roomId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, bookings)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings := []model.Booking{}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Find(&bookings).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error finding all bookings: %s", err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, bookings)
}

func GetBookingById(w http.ResponseWriter, r *http.Request) {
	booking := &model.Booking{}

	bookingId := chi.URLParam(r, "bookingID")
	bookingIdInt, err := strconv.Atoi(bookingId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&booking, bookingIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error finding booking with id %s: %s", bookingId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, booking)
}

func UpdateBookingById(w http.ResponseWriter, r *http.Request) {
	booking := &model.Booking{}
	booking_update := &model.Booking{}

	bookingId := chi.URLParam(r, "bookingId")
	bookingIdInt, err := strconv.Atoi(bookingId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking_update); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&booking, bookingIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error finding booking with id %s: %s", bookingId, err.Error()))
		return
	}

	if booking_update.CheckInDate != booking.CheckInDate {
		booking.CheckInDate = booking_update.CheckInDate
	}

	if booking_update.CheckOutDate != booking.CheckOutDate {
		booking.CheckOutDate = booking_update.CheckOutDate
	}

	if booking_update.TotalCost > 0 {
		booking.TotalCost = booking_update.TotalCost
	}

	if err := db.Save(&booking).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating booking with id %s: %s", bookingId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, booking)
}

func DelBookingById(w http.ResponseWriter, r *http.Request) {
	booking := &model.Booking{}

	bookingId := chi.URLParam(r, "bookingId")
	bookingIdInt, err := strconv.Atoi(bookingId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Delete(&booking, bookingIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting booking with id %s: %s", bookingId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, fmt.Sprintf("Successfully deleted booking with id %s", bookingId))
}
