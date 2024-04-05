package web

import (
	"encoding/json"
	"net/http"

	"github.com/beriloqueiroz/desafio-clean-arch/internal/entity"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/usecase"
	"github.com/beriloqueiroz/desafio-clean-arch/pkg/events"
)

type WebListOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderListedEvent events.EventInterface
}

func NewWebListOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderListedEvent events.EventInterface,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		OrderListedEvent: OrderListedEvent,
	}
}

func (h *WebListOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var dto usecase.ListOrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listOrder := usecase.NewListOrderUseCase(h.OrderRepository, h.OrderListedEvent, h.EventDispatcher)
	output, err := listOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
