package repository

import (
	"context"
	"fmt"
	"insightGlobal_carInventory/internal/infrastructure"
	"insightGlobal_carInventory/internal/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CarStorage interface {
	Create(ctx context.Context, car *model.Car) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Car, error)
	List(ctx context.Context, pageSize int, pageNumber int) ([]*model.Car, error)
	Update(ctx context.Context, car *model.Car) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type carStorage struct {
	dbConn infrastructure.DB
	log    *zap.SugaredLogger
}

func newCarStorage(db infrastructure.DB, logger *zap.SugaredLogger) CarStorage {
	return &carStorage{
		dbConn: db,
		log:    logger,
	}
}

func (c carStorage) Create(ctx context.Context, car *model.Car) error {
	_, err := c.dbConn.Insert().Model(car).Exec(ctx)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func (c carStorage) GetByID(ctx context.Context, id uuid.UUID) (*model.Car, error) {
	car := &model.Car{}
	err := c.dbConn.Select().Model(car).
		Where(fmt.Sprintf("%s = ?", model.StorageIdCar), id.String()).
		Where(fmt.Sprintf("%s = ?", model.StorageDisabledCar), false).
		Scan(ctx)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}

	return car, nil
}

func (c carStorage) List(ctx context.Context, pageSize int, pageNumber int) ([]*model.Car, error) {
	cars := make([]*model.Car, 0, pageSize)
	err := c.dbConn.Select().Model(&cars).
		Where(fmt.Sprintf("%s = ?", model.StorageDisabledCar), false).
		Limit(pageSize).Offset((pageNumber - 1) * pageSize).Scan(ctx)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}

	return cars, nil
}

func (c carStorage) Update(ctx context.Context, car *model.Car) error {
	_, err := c.dbConn.Update().Model(car).
		Where(fmt.Sprintf("%s = ?", model.StorageIdCar), car.ID.String()).
		Set(fmt.Sprintf("%s = ?", model.StorageDisabledCar), car.Disabled).
		Set(fmt.Sprintf("%s = ?", model.StorageColor), car.Color).
		Set(fmt.Sprintf("%s = ?", model.StorageMileage), car.Mileage).
		Set(fmt.Sprintf("%s = ?", model.StoragePrice), car.Price).
		Set(fmt.Sprintf("%s = NOW()", model.StorageUpd)).
		Where(fmt.Sprintf("%v = ?", model.StorageDisabledCar), false).
		Exec(ctx)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func (c carStorage) Delete(ctx context.Context, id uuid.UUID) error {
	var updated *model.Car
	_, err := c.dbConn.Update().
		Model(updated).
		Where(fmt.Sprintf("%s = ?", model.StorageIdCar), id.String()).
		Set(fmt.Sprintf("%s = ?", model.StorageDisabledCar), true).
		Exec(ctx)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}
