package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yogish-git/JeeraApp_Project/pkg/config"
)

var db *gorm.DB

type User struct {
	ID       uint `gorm:"primary_key"`
	USERNAME string
	Password string `gorm:"not null"`
	MAIL_ID  string
}

type Sprint struct {
	ID        uint `gorm:"primary_key"`
	NAME      string
	StartDate time.Time
	EndDate   time.Time
}

type Project struct {
	ID           uint `gorm:"primary_key"`
	Name         string
	TeamName     string
	ProjectOwner string
}

type UserStory struct {
	UST_ID             uint `gorm:"primary_key"`
	Title              string
	ProjectID          uint
	Project            Project `gorm:"foreignkey:ProjectID"`
	Reporter           string
	AcceptanceCriteria string
	Assignee           uint
	Date               time.Time
	Description        string
	EPIC_ID            uint
	EstimatedEfforts   int
	SprintID           uint
	Status             string
}

type EffortLogging struct {
	ID         uint `gorm:"primary_key"`
	TaskID     uint
	HoursSpent int
}

type Task struct {
	ID       uint `gorm:"primary_key"`
	UST_ID   uint
	Title    string
	Reporter string
	Assignee uint
	Status   string
}

type EPIC struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
}

type Issue struct {
	JIRA_ID            uint `gorm:"primary_key"`
	Title              string
	ProjectID          uint
	Project            Project `gorm:"foreignkey:ProjectID"`
	Reporter           string
	AcceptanceCriteria string
	Assignee           uint
	Date               time.Time
	Type               string
	Description        string
	EPIC_ID            uint
	EstimatedEfforts   int
	SprintID           uint
	Status             string
}

func init() {
	config.Connect()
	db = config.GetDB()
	// var DB = db
	// Auto-migrate the tables
	db.AutoMigrate(&User{}, &Sprint{}, &Project{}, &UserStory{}, &EffortLogging{}, &Task{}, &EPIC{}, &Issue{})

	// Define foreign key relationships
	db.Model(&UserStory{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Model(&UserStory{}).AddForeignKey("assignee", "users(id)", "CASCADE", "CASCADE")
	db.Model(&UserStory{}).AddForeignKey("sprint_id", "sprints(id)", "CASCADE", "CASCADE")
	db.Model(&UserStory{}).AddForeignKey("epic_id", "epics(id)", "CASCADE", "CASCADE")

	db.Model(&EffortLogging{}).AddForeignKey("task_id", "tasks(id)", "CASCADE", "CASCADE")

	db.Model(&Task{}).AddForeignKey("ust_id", "user_stories(ust_id)", "CASCADE", "CASCADE")
	db.Model(&Task{}).AddForeignKey("reporter", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Task{}).AddForeignKey("assignee", "users(id)", "CASCADE", "CASCADE")

	db.Model(&Issue{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Model(&Issue{}).AddForeignKey("assignee", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Issue{}).AddForeignKey("sprint_id", "sprints(id)", "CASCADE", "CASCADE")
	db.Model(&Issue{}).AddForeignKey("epic_id", "epics(id)", "CASCADE", "CASCADE")

	fmt.Println("Tables created and relationships established successfully!")
}

// database functions related to project table
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
