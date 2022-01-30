package usecase

import (
	"diary-app-service/entity"

	"github.com/stretchr/testify/mock"
)

type DiaryUseCaseMock struct {
	mock.Mock
}

func (m *DiaryUseCaseMock) CreateDiary(userID int, input *entity.CreateDiaryInput) (*entity.Diary, error) {
	args := m.Called(userID, input)

	return args.Get(0).(*entity.Diary), args.Error(1)
}

func (m *DiaryUseCaseMock) UpdateDiary(userID int, id int, input *entity.UpdateDiaryInput) (*entity.Diary, error) {
	args := m.Called(userID, id, input)

	return args.Get(0).(*entity.Diary), args.Error(1)
}

func (m *DiaryUseCaseMock) DeleteDiary(userID int, id int) (*entity.Diary, error) {
	args := m.Called(userID, id)

	return args.Get(0).(*entity.Diary), args.Error(1)
}

func (m *DiaryUseCaseMock) ListDiary(userID int, page int, pageSize int) (*[]entity.Diary, error) {
	args := m.Called(userID, page, pageSize)

	return args.Get(0).(*[]entity.Diary), args.Error(1)
}
