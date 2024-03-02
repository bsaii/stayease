package handler

import (
	"net/http"

	"github.com/bsaii/stayease/internal/utils"
)

func Error(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong.")
}
