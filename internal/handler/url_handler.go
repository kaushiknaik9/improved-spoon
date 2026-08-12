package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaushiknaik9/improved-spoon/internal/dto"
	"github.com/kaushiknaik9/improved-spoon/internal/response"
	"github.com/kaushiknaik9/improved-spoon/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

func (h *URLHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateURLRequest

	err := response.DecodeJSON(r, req)
	if err != nil {
		response.ErrorResponse(w, "invalid json received", http.StatusBadRequest)
		return
	}

	urlRecord, err := h.service.CreateShortURL(r.Context(), req.URL)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto.CreateURLResponse{
		ShortCode:   urlRecord.ShortCode,
		OriginalURL: urlRecord.OriginalURL,
	}

	response.JsonResponse(w, http.StatusOK, res)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	ShortCode := chi.URLParam(r, "shortCode")
	urlRecord, err := h.service.GetOriginalURL(r.Context(), ShortCode)
	if err != nil {
		response.ErrorResponse(w, "URL not found / INTERNALSERVERERROR", http.StatusBadRequest)
	}

	http.Redirect(w, r, urlRecord.OriginalURL, http.StatusFound)
}
