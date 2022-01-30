package repository

import "diary-app-service/entity"

type DiaryRepositoryInterface interface {
	GetDiary(id int, diary *entity.Diary) error
	CreateDiary(diary *entity.Diary) error
	UpdateDiary(id int, diary *entity.Diary, newDiary *entity.Diary) error
	DeleteDiary(id int, diary *entity.Diary) error
	ListDiary(user_id int, page int, pageSize int, diary *[]entity.Diary) error
}
