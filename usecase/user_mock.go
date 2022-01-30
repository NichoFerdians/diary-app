package usecase

import (
	"diary-app-service/entity"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (m *UserUseCaseMock) CreateUser(input *entity.UserInput) (*entity.User, error) {
	args := m.Called(input)

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *UserUseCaseMock) LoginUser(username string, password string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}
