package service

import (
	"context"
	"insightGlobal_carInventory/internal/model"
	"insightGlobal_carInventory/internal/repository"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type CarService interface {
	Create(ctx context.Context, car *model.Car) (*model.Car, error)
	GetByID(ctx context.Context, id string) (*model.Car, error)
	List(ctx context.Context, pageSize int, pageNumber int) ([]*model.Car, error)
	Update(ctx context.Context, car *model.Car) (*model.Car, error)
	Delete(ctx context.Context, id string) error
}

type carService struct {
	repository repository.Repository
	log        *zap.SugaredLogger
}

func (c carService) Create(ctx context.Context, car *model.Car) (*model.Car, error) {
	c.log.Info("Creating car with details:", car)
	car.ID = uuid.New()
	if err := c.repository.CarStorage().Create(ctx, car); err != nil {
		return nil, err
	}
	return car, nil
}

func (c carService) GetByID(ctx context.Context, id string) (*model.Car, error) {
	c.log.Info("Fetching car with ID:", id)
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid UUID: %s", id)
	}

	return c.repository.CarStorage().GetByID(ctx, uuid)
}

func (c carService) List(ctx context.Context, pageSize int, pageNumber int) ([]*model.Car, error) {
	c.log.Info("Fetching cars with page size:", pageSize, "and page number:", pageNumber)
	return c.repository.CarStorage().List(ctx, pageSize, pageNumber)
}

func (c carService) Update(ctx context.Context, car *model.Car) (*model.Car, error) {
	c.log.Info("Fetching car with ID:", car.ID)
	if existingCar, err := c.repository.CarStorage().GetByID(ctx, car.ID); err != nil {
		return nil, err
	} else if existingCar == nil {
		return nil, errors.Errorf("car with ID %s not found", car.ID)
	}

	if err := c.repository.CarStorage().Update(ctx, car); err != nil {
		return nil, err
	}

	//Return updated car
	return c.repository.CarStorage().GetByID(ctx, car.ID)
}

func (c carService) Delete(ctx context.Context, id string) error {
	c.log.Info("Fetching car with ID:", id)
	uuid, err := uuid.Parse(id)
	if err != nil {
		return errors.Wrapf(err, "invalid UUID: %s", id)
	}

	if existingCar, err := c.repository.CarStorage().GetByID(ctx, uuid); err != nil {
		return err
	} else if existingCar == nil {
		return errors.Errorf("car with ID %s not found", uuid)
	}

	return c.repository.CarStorage().Delete(ctx, uuid)
}

func NewCarService(r repository.Repository, l *zap.SugaredLogger) *carService {
	return &carService{
		repository: r,
		log:        l,
	}
}
