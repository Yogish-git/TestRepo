package models

import "github.com/jinzhu/gorm"

func (i *Task) CreateTask() *gorm.DB {
	db.NewRecord(i)
	result := db.Create(&i)
	return result
}

func GetTaskById(id int64) (*Task, *gorm.DB) {
	var requiredTask Task
	db := db.Where("id=?", id).Find(&requiredTask)
	return &requiredTask, db
}

func GetAllTask() []Task {
	var loggedTasks []Task
	db.Find(&loggedTasks)
	return loggedTasks
}

func DeleteTask(id int64) Task {
	var deletedTask Task
	db.Where("id=?", id).Delete(deletedTask)
	return deletedTask
}
