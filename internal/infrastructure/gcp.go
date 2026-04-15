package infrastructure

import (
	"context"
	"insightGlobal_carInventory/internal/config"
	"log"

	"cloud.google.com/go/storage"
)

func NewGCPConnection(ctx context.Context, gcp *config.GCP) {
	_, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("Failed to create GCP client: ", err)
	}

	log.Fatal("this is a test, it is not implemented yet")

	return
}
