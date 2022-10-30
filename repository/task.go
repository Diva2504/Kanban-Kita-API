package repository

import (
	"kanban-kita-api/models"

	"gorm.io/gorm"
)

func GetAllTask(db *gorm.DB, id uint) ([]models.Task, error) {
	var task []models.Task

	res := db.Where("user_id = ?", id).Find(&task)

	if res.Error != nil {
		return nil, res.Error
	} else {
		if res.RowsAffected <= 0 {
			return nil, res.Error
		} else {
			return task, res.Error
		}
	}
}

func GetDetail(db *gorm.DB, ID uint, userID uint) (models.Task, error) {
	task := models.Task{}

	err := db.Where("id = ? AND user_id = ?", ID, userID).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func CreateTask(data models.Task, db *gorm.DB) error {
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateTask(id int, userID uint, data models.Task, db *gorm.DB) (models.Task, error) {
	var task models.Task
	err := db.Model(&task).Where("id = ? AND user_id = ?", id, userID).Updates(&data).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, err
}

func DeleteTask(id int, db *gorm.DB) error {
	var task models.Task

	del := db.Delete(&task, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}
}

func UpdateCategoryTask(db *gorm.DB, id int, categoryId uint) error {
	var task models.Task
	err := db.Model(&task).Where("id = ?", id).Updates(models.Task{CategoryId: categoryId}).Error
	return err
}
