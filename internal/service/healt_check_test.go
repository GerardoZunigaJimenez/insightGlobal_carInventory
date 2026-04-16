package service

import (
	"context"
	"errors"
	infraMocks "insightGlobal_carInventory/internal/infrastructure/mocks"
	repository "insightGlobal_carInventory/internal/repository"
	repoMocks "insightGlobal_carInventory/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckService_PingMysql(t *testing.T) {

	ctx := context.Background()
	var mockRepo *repoMocks.MockRepository
	var mockInfra *infraMocks.MockDB

	tests := []struct {
		name          string
		repository    func() repository.Repository
		repositoryErr error
		wantErr       error
	}{
		{
			name: "PingMysql",
			repository: func() repository.Repository {
				mockRepo = repoMocks.NewMockRepository(t)
				mockInfra = nil
				return mockRepo
			},
		},
		{
			name: "PingMysql with error",
			repository: func() repository.Repository {
				mockRepo = repoMocks.NewMockRepository(t)
				mockInfra = nil
				return mockRepo
			},
			repositoryErr: errors.New("expected error"),
			wantErr:       errors.New("expected error"),
		},
		{
			name: "unable to ping the DB, because the reference is nil",
			repository: func() repository.Repository {
				dbConn := infraMocks.NewMockDB(t)
				r := repository.NewRepositories(dbConn, nil)
				mockInfra = dbConn
				mockRepo = nil
				return r
			},
			wantErr: errors.New("unable to ping the DB, because the reference is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewHealthCheckService(tt.repository())

			if mockRepo != nil {
				mockRepo.On("Ping", ctx).Return(tt.repositoryErr)
			}
			if mockInfra != nil {
				mockInfra.On("IsAlive").Return(errors.New("db connection has not been set up"))
			}

			err := s.PingMysql(ctx)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}
