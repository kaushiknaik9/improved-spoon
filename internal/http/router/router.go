package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaushiknaik9/improved-spoon/internal/handler"
)

type URLHandler struct {
	handler *handler.URLHandler
}

func NewURLHandler(handler *handler.URLHandler) *URLHandler {
	return &URLHandler{
		handler: handler,
	}
}

func (h *URLHandler) RegisterURLRoutes(r chi.Router) {
	r.Post("/shorten", h.handler.CreateShortUrl)
	r.Get("/{shortCode}", h.handler.Redirect)
}
