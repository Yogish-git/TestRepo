package models

import "github.com/jinzhu/gorm"

//database functions related to effortLogging
func (u *EffortLogging) CreateEffort() *gorm.DB {
	db.NewRecord(u)
	result := db.Create(&u)

	return result
}

func GetEffortById(id int64) (*EffortLogging, *gorm.DB) {
	var requiredEffort EffortLogging
	db := db.Where("id=?", id).Find(&requiredEffort)
	return &requiredEffort, db
}

func Getalleffort() []EffortLogging {
	var loggedEfforts []EffortLogging
	db.Find(&loggedEfforts)
	return loggedEfforts
}

func DeleteEffort(id int64) EffortLogging {
	var deletedEffort EffortLogging
	db.Where("id=?", id).Delete(deletedEffort)
	return deletedEffort
}
