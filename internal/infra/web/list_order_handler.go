package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	pageInt := 0
	pageSizeInt := 1000
	var err error = nil
	if len(r.URL.Query().Get("page")) > 0 && len(r.URL.Query().Get("pageSize")) > 0 {
		pageInt, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageInt = 0
		}
		pageSizeInt, err = strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil {
			pageInt = 0
			pageSizeInt = 1000
		}
	}

	dto := usecase.ListOrderInputDTO{
		Page:     pageInt,
		PageSize: pageSizeInt,
	}
	fmt.Println("listing order...")
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
