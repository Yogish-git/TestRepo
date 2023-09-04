package models

import "github.com/jinzhu/gorm"

//database functions related to userstories
func (u *UserStory) CreateUST() *gorm.DB {
	db.NewRecord(u)
	result := db.Create(&u)

	return result
}

func GetUSTById(id int64) (*UserStory, *gorm.DB) {
	var requiredUST UserStory
	db := db.Where("UST_ID=?", id).Find(&requiredUST)
	return &requiredUST, db
}

func Getallust() []UserStory {
	var UserStories []UserStory
	db.Find(&UserStories)
	return UserStories
}

func DeleteUST(id int64) UserStory {
	var deletedUST UserStory
	db.Where("UST_ID=?", id).Delete(deletedUST)
	return deletedUST
}
