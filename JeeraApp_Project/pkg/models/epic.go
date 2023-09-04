package models

import "github.com/jinzhu/gorm"

//database functions related to EpicLogging
func (u *EPIC) CreateEpic() *gorm.DB {
	db.NewRecord(u)
	result := db.Create(&u)

	return result
}

func GetEpicById(id int64) (*EPIC, *gorm.DB) {
	var requiredEpic EPIC
	db := db.Where("id=?", id).Find(&requiredEpic)
	return &requiredEpic, db
}

func GetAllEpic() []EPIC {
	var loggedEpics []EPIC
	db.Find(&loggedEpics)
	return loggedEpics
}

func DeleteEpic(id int64) EPIC {
	var deletedEpic EPIC
	db.Where("id=?", id).Delete(deletedEpic)
	return deletedEpic
}
