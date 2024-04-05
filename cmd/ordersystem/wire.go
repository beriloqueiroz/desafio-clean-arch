//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/beriloqueiroz/desafio-clean-arch/internal/entity"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/event"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/database"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/web"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/usecase"
	"github.com/beriloqueiroz/desafio-clean-arch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrderListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

// list

var setOrderListedEvent = wire.NewSet(
	event.NewOrderListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
)

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebListOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebListOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		web.NewWebListOrderHandler,
	)
	return &web.WebListOrderHandler{}
}
