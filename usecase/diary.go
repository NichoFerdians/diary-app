package usecase

import (
	"diary-app-service/entity"
	"diary-app-service/repository"
	"errors"
)

type DiaryUseCaseInterface interface {
	CreateDiary(userID int, input *entity.CreateDiaryInput) (*entity.Diary, error)
	UpdateDiary(userID int, id int, input *entity.UpdateDiaryInput) (*entity.Diary, error)
	DeleteDiary(userID int, id int) (*entity.Diary, error)
	ListDiary(userID int, page int, pageSize int) (*[]entity.Diary, error)
}

type DiaryUseCase struct {
	diaryRepo repository.DiaryRepositoryInterface
}

func NewDiaryUseCase(
	diaryRepo repository.DiaryRepositoryInterface) *DiaryUseCase {
	return &DiaryUseCase{
		diaryRepo: diaryRepo,
	}
}

func (u DiaryUseCase) CreateDiary(userID int, input *entity.CreateDiaryInput) (*entity.Diary, error) {

	newDairy := entity.Diary{
		Title:  input.Title,
		Body:   input.Body,
		UserID: userID,
	}

	err := u.diaryRepo.CreateDiary(&newDairy)
	if err != nil {
		return nil, err
	}

	return &newDairy, nil
}

func (u DiaryUseCase) UpdateDiary(userID int, id int, input *entity.UpdateDiaryInput) (*entity.Diary, error) {
	var diary entity.Diary

	newDairy := entity.Diary{
		Title: input.Title,
		Body:  input.Body,
	}

	err := u.diaryRepo.GetDiary(id, &diary)
	if err != nil {
		return nil, err
	}

	if diary.UserID != userID {
		return nil, errors.New("diary doesn't belong to the authenticated user")
	}

	err = u.diaryRepo.UpdateDiary(id, &diary, &newDairy)
	if err != nil {
		return nil, err
	}

	return &diary, nil
}

func (u DiaryUseCase) DeleteDiary(userID int, id int) (*entity.Diary, error) {
	var diary entity.Diary

	err := u.diaryRepo.GetDiary(id, &diary)
	if err != nil {
		return nil, err
	}

	if diary.UserID != userID {
		return nil, errors.New("diary doesn't belong to the authenticated user")
	}

	err = u.diaryRepo.DeleteDiary(id, &diary)
	if err != nil {
		return nil, err
	}

	return &diary, nil
}

func (u DiaryUseCase) ListDiary(userID int, page int, pageSize int) (*[]entity.Diary, error) {
	var diary []entity.Diary

	err := u.diaryRepo.ListDiary(userID, page, pageSize, &diary)
	if err != nil {
		return nil, err
	}

	return &diary, nil
}
