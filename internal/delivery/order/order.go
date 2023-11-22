package order

import (
	"context"
	"order/internal/repo/postgres"
	"order/internal/types"
	"order/pkg/order/pkg/prots"
)

type OrderHandler struct {
	prots.UnimplementedOrderServiceServer
	r *postgres.Repo
}

func NewOrder(r *postgres.Repo) *OrderHandler {
	return &OrderHandler{
		r: r,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *prots.CreateOrderRequest) (*prots.CreateOrderResponse, error) {
	payload := types.CreateOrderUUID{
		UUID:  req.UserUuid,
		Order: types.Menu{},
	}
	if err := h.r.CreateOrder(ctx, payload); err != nil {
		return new(prots.CreateOrderResponse), err
	}
	return new(prots.CreateOrderResponse), nil
}

func (h *OrderHandler) GetActualMenu(ctx context.Context, req *prots.GetActualMenuRequest) (*prots.GetActualMenuResponse, error) {
	_, err := h.r.GetMenu(ctx)
	if err != nil {
		return new(prots.GetActualMenuResponse), err
	}
	return new(prots.GetActualMenuResponse), nil
}
