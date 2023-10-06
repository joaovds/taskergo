package repositories

import (
  "github.com/joaovds/taskergo/internal/domain/entities"
)

type LoadTaskGroupsRepository interface {
  LoadAll() ([]LoadTaskGroupRepositoryResult, error)
}

type LoadTaskGroupRepositoryResult struct {
  entities.TaskGroup
}

func (l LoadTaskGroupRepositoryResult) MapToTaskGroup() entities.TaskGroup {
  return l.TaskGroup
}

