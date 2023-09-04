package models

import "github.com/jinzhu/gorm"

func (i *Issue) CreateIssue() *gorm.DB {
	db.NewRecord(i)
	result := db.Create(&i)
	return result
}

func GetIssueById(id int64) (*Issue, *gorm.DB) {
	var requiredIssue Issue
	db := db.Where("id=?", id).Find(&requiredIssue)
	return &requiredIssue, db
}

func GetAllIssue() []Issue {
	var loggedIssues []Issue
	db.Find(&loggedIssues)
	return loggedIssues
}

func DeleteIssue(id int64) Issue {
	var deletedIssue Issue
	db.Where("id=?", id).Delete(deletedIssue)
	return deletedIssue
}
