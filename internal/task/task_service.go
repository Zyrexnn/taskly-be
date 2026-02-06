package task

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateTask(userID uint, req CreateTaskDTO) (*Task, error)
	GetAllTasks(userID uint) ([]Task, error)
	GetTaskByID(userID, taskID uint) (*Task, error)
	UpdateTask(userID, taskID uint, req UpdateTaskDTO) (*Task, error)
	DeleteTask(userID, taskID uint) error
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s *service) CreateTask(userID uint, req CreateTaskDTO) (*Task, error) {
	task := Task{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
	}
	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *service) GetAllTasks(userID uint) ([]Task, error) {
	var tasks []Task
	if err := s.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *service) GetTaskByID(userID, taskID uint) (*Task, error) {
	var task Task
	if err := s.db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (s *service) UpdateTask(userID, taskID uint, req UpdateTaskDTO) (*Task, error) {
	task, err := s.GetTaskByID(userID, taskID)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (s *service) DeleteTask(userID, taskID uint) error {
	task, err := s.GetTaskByID(userID, taskID)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}