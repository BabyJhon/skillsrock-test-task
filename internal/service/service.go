package service

import (
	"context"

	"github.com/BabyJhon/skillsrock-test-task/internal/entity"
	"github.com/BabyJhon/skillsrock-test-task/internal/repository"
)

type Task interface {
	CreateTask(ctx context.Context, task entity.Task) (int, error)
	DeleteTask(ctx context.Context, id int) (int, error)
	UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	GetAllTasks(ctx context.Context) (*[]entity.Task, error)
}

type Service struct {
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repos),
	}
}
