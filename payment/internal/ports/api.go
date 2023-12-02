package ports

import (
	"context"
	"github.com/w3gop2p/microservices/payment/internal/appplication/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
