package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	StorageDisabledCar = "disabled"
	StorageColor       = "color"
	StorageMileage     = "mileage"
	StoragePrice       = "price"
	StorageUpd         = "updated_at"
)

type Car struct {
	ID        uuid.UUID `json:"id" bun:"id,pk,type:uuid"`
	Make      string    `json:"make" bun:"make"`
	Model     string    `json:"model" bun:"model"`
	Year      int       `json:"year" bun:"year"`
	Color     string    `json:"color" bun:"color"`
	VIN       string    `json:"vin" bun:"vin"`
	Mileage   int       `json:"mileage" bun:"mileage"`
	Price     float64   `json:"price" bun:"price"`
	Disabled  bool      `json:"-" bun:"disabled"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at"`
}
