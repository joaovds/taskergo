package repositories

import "github.com/joaovds/taskergo/internal/domain/entities"

type LoadTaskGroupsRepository interface {
  LoadAll() ([]LoadTaskGroupRepositoryResult, error)
}

type LoadTaskGroupRepositoryResult struct {
  entities.TaskGroup
}

// função para fazer um map do retorno de loadALl para o load do usecase do domain
func (l LoadTaskGroupRepositoryResult) MapToTaskGroup() entities.TaskGroup {
  return l.TaskGroup
}

