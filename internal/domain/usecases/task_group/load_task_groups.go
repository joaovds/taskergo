package usecases

import "github.com/joaovds/taskergo/internal/domain/entities"

type LoadTaskGroups interface {
  Load() ([]entities.TaskGroup, error)
}

