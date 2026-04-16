package service

import (
	"context"
	"errors"
	"insightGlobal_carInventory/internal/model"
	repositoryMocks "insightGlobal_carInventory/internal/repository/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestCarService_Create(t *testing.T) {
	ctx := context.Background()
	log := zap.Must(zap.NewDevelopment()).Sugar()

	tests := map[string]struct {
		inputCar   *model.Car
		mockSetup  func(*repositoryMocks.MockRepository, *repositoryMocks.MockCarStorage, *model.Car)
		assertFunc func(t *testing.T, got *model.Car, err error, inputCar *model.Car)
	}{
		"success": {
			inputCar: &model.Car{
				Make:  "Toyota",
				Model: "Corolla",
				Year:  2020,
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("Create", ctx, inputCar).Return(nil)
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.NotEqual(t, "", got.ID.String())
				assert.Equal(t, inputCar.Make, got.Make)
				assert.Equal(t, inputCar.Model, got.Model)
				assert.Equal(t, inputCar.Year, got.Year)
			},
		},
		"repository returns error": {
			inputCar: &model.Car{
				Make:  "Honda",
				Model: "Civic",
				Year:  2021,
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("Create", ctx, inputCar).Return(errors.New("create failed"))
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.Error(t, err)
				assert.EqualError(t, err, "create failed")
				assert.Nil(t, got)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repositoryMocks.NewMockRepository(t)
			carStorage := repositoryMocks.NewMockCarStorage(t)

			test.mockSetup(repo, carStorage, test.inputCar)

			s := NewCarService(repo, log)
			got, err := s.Create(ctx, test.inputCar)

			test.assertFunc(t, got, err, test.inputCar)

			repo.AssertExpectations(t)
			carStorage.AssertExpectations(t)
		})
	}
}

func TestCarService_GetByID(t *testing.T) {
	ctx := context.Background()
	log := zap.Must(zap.NewDevelopment()).Sugar()

	tests := map[string]struct {
		id         string
		mockSetup  func(*repositoryMocks.MockRepository, *repositoryMocks.MockCarStorage, uuid.UUID)
		assertFunc func(t *testing.T, got *model.Car, err error)
	}{
		"success": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				expectedCar := &model.Car{
					ID:    id,
					Make:  "Toyota",
					Model: "Corolla",
					Year:  2020,
				}
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return(expectedCar, nil)
			},
			assertFunc: func(t *testing.T, got *model.Car, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.Equal(t, "Toyota", got.Make)
				assert.Equal(t, "Corolla", got.Model)
				assert.Equal(t, 2020, got.Year)
			},
		},
		"repository returns error": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return((*model.Car)(nil), errors.New("error to get car"))
			},
			assertFunc: func(t *testing.T, got *model.Car, err error) {
				require.Error(t, err)
				assert.EqualError(t, err, "error to get car")
				assert.Nil(t, got)
			},
		},
		"invalid uuid": {
			id: "invalid-uuid",
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				// No mock setup needed because UUID parsing fails before repository call.
			},
			assertFunc: func(t *testing.T, got *model.Car, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "invalid UUID")
				assert.Nil(t, got)
			},
		},
		"not found": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				var expectedCar *model.Car = nil
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return(expectedCar, nil)
			},
			assertFunc: func(t *testing.T, got *model.Car, err error) {
				require.NoError(t, err)
				require.Nil(t, got)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repositoryMocks.NewMockRepository(t)
			carStorage := repositoryMocks.NewMockCarStorage(t)

			if name != "invalid uuid" {
				parsedID, err := uuid.Parse(test.id)
				require.NoError(t, err)
				test.mockSetup(repo, carStorage, parsedID)
			}

			s := NewCarService(repo, log)
			got, err := s.GetByID(ctx, test.id)

			test.assertFunc(t, got, err)

			repo.AssertExpectations(t)
			carStorage.AssertExpectations(t)
		})
	}
}

func TestCarService_List(t *testing.T) {
	ctx := context.Background()
	log := zap.Must(zap.NewDevelopment()).Sugar()

	tests := map[string]struct {
		pageSize   int
		pageNumber int
		mockSetup  func(*repositoryMocks.MockRepository, *repositoryMocks.MockCarStorage, int, int)
		assertFunc func(t *testing.T, got []*model.Car, err error)
	}{
		"success": {
			pageSize:   10,
			pageNumber: 1,
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, pageSize int, pageNumber int) {
				expectedCars := []*model.Car{
					{
						ID:    uuid.New(),
						Make:  "Toyota",
						Model: "Corolla",
						Year:  2020,
					},
					{
						ID:    uuid.New(),
						Make:  "Honda",
						Model: "Civic",
						Year:  2021,
					},
				}
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("List", ctx, pageSize, pageNumber).Return(expectedCars, nil)
			},
			assertFunc: func(t *testing.T, got []*model.Car, err error) {
				require.NoError(t, err)
				require.Len(t, got, 2)
				assert.Equal(t, "Toyota", got[0].Make)
				assert.Equal(t, "Corolla", got[0].Model)
				assert.Equal(t, "Honda", got[1].Make)
				assert.Equal(t, "Civic", got[1].Model)
			},
		},
		"repository returns error": {
			pageSize:   20,
			pageNumber: 2,
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, pageSize int, pageNumber int) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("List", ctx, pageSize, pageNumber).Return(([]*model.Car)(nil), errors.New("error listing cars"))
			},
			assertFunc: func(t *testing.T, got []*model.Car, err error) {
				require.Error(t, err)
				assert.EqualError(t, err, "error listing cars")
				assert.Nil(t, got)
			},
		},
		"empty result": {
			pageSize:   5,
			pageNumber: 3,
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, pageSize int, pageNumber int) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("List", ctx, pageSize, pageNumber).Return([]*model.Car{}, nil)
			},
			assertFunc: func(t *testing.T, got []*model.Car, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.Empty(t, got)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repositoryMocks.NewMockRepository(t)
			carStorage := repositoryMocks.NewMockCarStorage(t)

			test.mockSetup(repo, carStorage, test.pageSize, test.pageNumber)

			s := NewCarService(repo, log)
			got, err := s.List(ctx, test.pageSize, test.pageNumber)

			test.assertFunc(t, got, err)

			repo.AssertExpectations(t)
			carStorage.AssertExpectations(t)
		})
	}
}

func TestCarService_Update(t *testing.T) {
	ctx := context.Background()
	log := zap.Must(zap.NewDevelopment()).Sugar()

	tests := map[string]struct {
		inputCar   *model.Car
		mockSetup  func(*repositoryMocks.MockRepository, *repositoryMocks.MockCarStorage, *model.Car)
		assertFunc func(t *testing.T, got *model.Car, err error, inputCar *model.Car)
	}{
		"success": {
			inputCar: &model.Car{
				ID:    uuid.New(),
				Make:  "Toyota",
				Model: "Corolla",
				Year:  2020,
				Color: "Red",
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				updatedCar := &model.Car{
					ID:    inputCar.ID,
					Make:  inputCar.Make,
					Model: inputCar.Model,
					Year:  inputCar.Year,
					Color: inputCar.Color,
				}
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, inputCar.ID).Return(inputCar, nil).Once()
				carStorage.On("Update", ctx, updatedCar).Return(nil).Once()
				carStorage.On("GetByID", ctx, inputCar.ID).Return(updatedCar, nil).Once()
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.Equal(t, inputCar.ID, got.ID)
				assert.Equal(t, inputCar.Make, got.Make)
				assert.Equal(t, inputCar.Model, got.Model)
				assert.Equal(t, inputCar.Year, got.Year)
				assert.Equal(t, inputCar.Color, got.Color)
			},
		},
		"not found": {
			inputCar: &model.Car{
				ID:    uuid.New(),
				Make:  "Honda",
				Model: "Civic",
				Year:  2021,
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, inputCar.ID).Return((*model.Car)(nil), nil)
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.Error(t, err)
				assert.EqualError(t, err, "car with ID "+inputCar.ID.String()+" not found")
				assert.Nil(t, got)
			},
		},
		"repository get by id returns error": {
			inputCar: &model.Car{
				ID:   uuid.New(),
				Make: "Ford",
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, inputCar.ID).Return((*model.Car)(nil), errors.New("get by id failed"))
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.Error(t, err)
				assert.EqualError(t, err, "get by id failed")
				assert.Nil(t, got)
			},
		},
		"repository update returns error": {
			inputCar: &model.Car{
				ID:    uuid.New(),
				Make:  "Nissan",
				Model: "Versa",
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, inputCar.ID).Return(inputCar, nil).Once()
				carStorage.On("Update", ctx, inputCar).Return(errors.New("update failed")).Once()
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.Error(t, err)
				assert.EqualError(t, err, "update failed")
				assert.Nil(t, got)
			},
		},
		"repository get updated car returns error": {
			inputCar: &model.Car{
				ID:    uuid.New(),
				Make:  "Mazda",
				Model: "CX-5",
			},
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, inputCar *model.Car) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, inputCar.ID).Return(inputCar, nil).Once()
				carStorage.On("Update", ctx, inputCar).Return(nil).Once()
				carStorage.On("GetByID", ctx, inputCar.ID).Return((*model.Car)(nil), errors.New("final get failed")).Once()
			},
			assertFunc: func(t *testing.T, got *model.Car, err error, inputCar *model.Car) {
				require.Error(t, err)
				assert.EqualError(t, err, "final get failed")
				assert.Nil(t, got)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repositoryMocks.NewMockRepository(t)
			carStorage := repositoryMocks.NewMockCarStorage(t)

			test.mockSetup(repo, carStorage, test.inputCar)

			s := NewCarService(repo, log)
			got, err := s.Update(ctx, test.inputCar)

			test.assertFunc(t, got, err, test.inputCar)

			repo.AssertExpectations(t)
			carStorage.AssertExpectations(t)
		})
	}
}

func TestCarService_Delete(t *testing.T) {
	ctx := context.Background()
	log := zap.Must(zap.NewDevelopment()).Sugar()

	tests := map[string]struct {
		id         string
		mockSetup  func(*repositoryMocks.MockRepository, *repositoryMocks.MockCarStorage, uuid.UUID)
		assertFunc func(t *testing.T, err error)
	}{
		"success": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				existingCar := &model.Car{
					ID:   id,
					Make: "Toyota",
				}
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return(existingCar, nil).Once()
				carStorage.On("Delete", ctx, id).Return(nil).Once()
			},
			assertFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"repository get by id returns error": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return((*model.Car)(nil), errors.New("get by id failed")).Once()
			},
			assertFunc: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.EqualError(t, err, "get by id failed")
			},
		},
		"invalid uuid": {
			id: "invalid-uuid",
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				// No mock setup needed because UUID parsing fails before repository call.
			},
			assertFunc: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "invalid UUID")
			},
		},
		"not found": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return((*model.Car)(nil), nil).Once()
			},
			assertFunc: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "not found")
			},
		},
		"delete returns error": {
			id: uuid.NewString(),
			mockSetup: func(repo *repositoryMocks.MockRepository, carStorage *repositoryMocks.MockCarStorage, id uuid.UUID) {
				existingCar := &model.Car{
					ID:   id,
					Make: "Honda",
				}
				repo.On("CarStorage").Return(carStorage)
				carStorage.On("GetByID", ctx, id).Return(existingCar, nil).Once()
				carStorage.On("Delete", ctx, id).Return(errors.New("delete failed")).Once()
			},
			assertFunc: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.EqualError(t, err, "delete failed")
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repositoryMocks.NewMockRepository(t)
			carStorage := repositoryMocks.NewMockCarStorage(t)

			if name != "invalid uuid" {
				parsedID, err := uuid.Parse(test.id)
				require.NoError(t, err)
				test.mockSetup(repo, carStorage, parsedID)
			}

			s := NewCarService(repo, log)
			err := s.Delete(ctx, test.id)

			test.assertFunc(t, err)

			repo.AssertExpectations(t)
			carStorage.AssertExpectations(t)
		})
	}
}
