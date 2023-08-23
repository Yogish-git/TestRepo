package models

import (
	"github.com/jinzhu/gorm"
	"github.com/yogish-git/JeeraApp_Project/pkg/config"
)

var db *gorm.DB

type Project struct {
	// gorm.Model
	Name string `gorm:""json:"name"`
	Id   int    `json:"id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Project{})
}

func (p *Project) CreateProject() *Project {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllProjects() []Project {
	var Projects []Project
	db.Find(&Projects)
	return Projects

}

func GetProjectById(Id int64) (*Project, *gorm.DB) {
	var getProject Project
	db := db.Where("Id=?", Id).Find(&getProject)
	return &getProject, db
}

func DeleteProject(Id int64) Project {
	var project Project
	db.Where("Id=?", Id).Delete(project)
	return project
}
