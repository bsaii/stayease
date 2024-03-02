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

func AddRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&room); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Validate the JSON data

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Create(&room).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating record for room: %s", err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, room)

}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}

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

	if err := db.First(&room, roomIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error finding room with id %s: %s", roomId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, room)
}

func AllRooms(w http.ResponseWriter, r *http.Request) {
	rooms := []model.Room{}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.Find(&rooms).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error finding all rooms: %s", err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, rooms)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}
	room_update := &model.Room{}

	roomId := chi.URLParam(r, "roomID")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter format")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&room_update); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database not found in context")
		return
	}

	if err := db.First(&room, roomIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error finding room with id %s: %s", roomId, err.Error()))
		return
	}

	if room_update.Capacity > 0 {
		room.Capacity = room_update.Capacity
	}
	if room_update.Description != "" {
		room.Description = room_update.Description
	}
	if room_update.Price > 0 {
		room.Price = room_update.Price
	}
	if room_update.RoomNumber != "" {
		room.RoomNumber = room_update.RoomNumber
	}
	if room_update.Type != "" {
		room.Type = room_update.Type
	}

	if err := db.Save(&room).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating room with id %s: %s", roomId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, room)
}

func DelRoom(w http.ResponseWriter, r *http.Request) {
	room := &model.Room{}

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

	if err := db.Delete(&room, roomIdInt).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting room with id %s: %s", roomId, err.Error()))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, fmt.Sprintf("Successfully deleted room with id %s", roomId))
}
