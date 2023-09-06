package models

import "github.com/jinzhu/gorm"

func (i *Sprint) CreateSprint() *gorm.DB {
	db.NewRecord(i)
	result := db.Create(&i)
	return result
}

func GetSprintById(id int64) (*Sprint, *gorm.DB) {
	var requiredSprint Sprint
	db := db.Where("id=?", id).Find(&requiredSprint)
	return &requiredSprint, db
}

func GetAllSprint() []Sprint {
	var loggedSprints []Sprint
	db.Find(&loggedSprints)
	return loggedSprints
}

func DeleteSprint(id int64) Sprint {
	var deletedSprint Sprint
	db.Where("id=?", id).Delete(deletedSprint)
	return deletedSprint
}
