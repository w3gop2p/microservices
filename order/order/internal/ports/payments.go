package ports

import (
	"context"
	"github.com/w3gop2p/microservices/order/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
