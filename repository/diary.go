package repository

import (
	"diary-app-service/entity"
	"diary-app-service/helper/pagination"
	"errors"

	"gorm.io/gorm"
)

type DiaryRepository struct {
	db *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) *DiaryRepository {
	return &DiaryRepository{
		db: db,
	}
}

func (r DiaryRepository) GetDiary(id int, diary *entity.Diary) error {

	if err := r.db.Where("id = ?", id).First(&diary).Error; err != nil {
		return errors.New("diary not found")
	}

	return nil
}

func (r DiaryRepository) CreateDiary(diary *entity.Diary) error {

	r.db.Create(&diary)

	return nil
}

func (r DiaryRepository) UpdateDiary(id int, diary *entity.Diary, newDiary *entity.Diary) error {

	r.db.Model(&diary).Updates(&newDiary)

	return nil
}

func (r DiaryRepository) DeleteDiary(id int, diary *entity.Diary) error {

	r.db.Delete(&diary)

	return nil
}

func (r DiaryRepository) ListDiary(user_id int, page int, pageSize int, diary *[]entity.Diary) error {

	r.db.Scopes(pagination.Paginate(page, pageSize)).Where("user_id = ?", user_id).Find(&diary)

	return nil
}
