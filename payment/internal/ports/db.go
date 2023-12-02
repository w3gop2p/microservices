package ports

import (
	"context"
	"github.com/w3gop2p/microservices/payment/internal/appplication/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}
