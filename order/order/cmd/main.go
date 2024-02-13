package main

import (
	"github.com/w3gop2p/microservices/order/order/config"
	"github.com/w3gop2p/microservices/order/order/internal/adapters/db"
	"github.com/w3gop2p/microservices/order/order/internal/adapters/grpc"
	"github.com/w3gop2p/microservices/order/order/internal/adapters/payment"
	"github.com/w3gop2p/microservices/order/order/internal/application/core/api"
	"log"
	"os"
)

func main() {
	err := os.Setenv("DATA_SOURCE_URL", "root:parola123@tcp(127.0.0.1:3306)/order123")
	if err != nil {
		return
	}
	err = os.Setenv("APPLICATION_PORT", "4000")
	if err != nil {
		return
	}
	err = os.Setenv("ENV", "development")
	if err != nil {
		return
	}
	err = os.Setenv("PAYMENT_SERVICE_URL", "localhost:4001")
	if err != nil {
		return
	}

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
