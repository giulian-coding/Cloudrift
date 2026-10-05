package order

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/giulian-coding/cloudrift/internal/platform/httpx"

	"uuid"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.placeOrder)
	mux.HandleFunc("GET /orders/{id}", h.getOrder)
}

func (h *Handler) placeOrder(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	ctx := r.Context()

	order, err := h.svc.PlaceOrder(ctx, req)
	if err != nil {
		if errors.Is(err, ErrCustomerRequired) {
			httpx.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusCreated, order)
}

func (h *Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	safeID, err := uuid.Parse(id)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	order, err := h.svc.Get(ctx, safeID.String())
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.Error(w, http.StatusNotFound, err.Error())
			return
		}
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, order)
}
