package ports

import (
	"context"
	"github.com/w3gop2p/microservices/order/order/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(context.Context, *domain.Order) error
}
