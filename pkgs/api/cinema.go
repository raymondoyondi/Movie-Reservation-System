package api

import (
	"net/http"

	jsonHelper "github.com/raymondoyondi/Movie-Reservation-System/internal/json"
	"github.com/raymondoyondi/Movie-Reservation-System/pkgs/service/cinema"
)

func (h *Handler) ListCinemas(request *http.Request, writer http.ResponseWriter) {
	if out, err := h.svc.cinema.ListCinemas(); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddCinema(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp cinema.AddCinemaInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	// process input and return output
	if out, err := h.svc.cinema.AddCinema(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddCinemaScreen(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp cinema.AddCinemaScreenInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	// process input and return output
	if out, err := h.svc.cinema.AddCinemaScreen(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}
