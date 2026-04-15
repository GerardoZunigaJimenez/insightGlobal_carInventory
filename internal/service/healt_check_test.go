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
	t.Parallel()
	ctx := context.Background()
	mockRepo := repoMocks.NewMockRepository(t)

	tests := []struct {
		name          string
		repository    repository.Repository
		repositoryErr error
		wantErr       error
	}{
		{
			name:       "PingMysql",
			repository: mockRepo,
		},
		{
			name:          "PingMysql",
			repository:    mockRepo,
			repositoryErr: errors.New("expected error"),
			wantErr:       errors.New("expected error"),
		},
		{
			name: "unable to ping the DB, because the reference is nil",
			repository: func() repository.Repository {
				dbConn := infraMocks.NewMockDB(t)
				dbConn.On("IsAlive").Return(errors.New("db connection has not been set up"))
				r := repository.NewRepositories(dbConn, nil)
				return r
			}(),
			wantErr: errors.New("unable to ping the DB, because the reference is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repository != nil {
				mockRepo.On("Ping", ctx).Return(tt.repositoryErr)
			}

			s := newHealthCheckService(tt.repository)
			err := s.PingMysql(ctx)
			assert.Equal(t, err, tt.wantErr)
		})
	}

}
