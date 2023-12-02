package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/w3gop2p/microservices/payment/adapters/db"
	"github.com/w3gop2p/microservices/payment/adapters/grpc"
	"github.com/w3gop2p/microservices/payment/config"
	"github.com/w3gop2p/microservices/payment/internal/appplication/core/api"
	"os"
)

const (
	service     = "payment"
	environment = "dev"
	id          = 2
)

func main() {
	err := os.Setenv("DATA_SOURCE_URL", "root:@tcp(127.0.0.1:3306)/payment21")
	if err != nil {
		return
	}
	err = os.Setenv("APPLICATION_PORT", "4001")
	if err != nil {
		return
	}
	err = os.Setenv("ENV", "development")
	if err != nil {
		return
	}

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
