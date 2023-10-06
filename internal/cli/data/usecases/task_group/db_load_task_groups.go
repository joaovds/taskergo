package data_usecases

import (
	repositories "github.com/joaovds/taskergo/internal/cli/data/contracts/repositories/task_group"
	"github.com/joaovds/taskergo/internal/domain/entities"
)

type DbLoadTaskGroupsUsecase struct {
  loadTaskGroupsRepository repositories.LoadTaskGroupsRepository
}

func DbLoadTaskGroups(loadTaskGroupsRepository repositories.LoadTaskGroupsRepository) *DbLoadTaskGroupsUsecase {
  return &DbLoadTaskGroupsUsecase{loadTaskGroupsRepository: loadTaskGroupsRepository}
}

func (d *DbLoadTaskGroupsUsecase) Load() ([]entities.TaskGroup, error) {
  taskGroups, err := d.loadTaskGroupsRepository.LoadAll()
  if err != nil {
    return nil, err
  }

  taskGroupMap := make([]entities.TaskGroup, len(taskGroups))

  for i, taskGroup := range taskGroups {
    taskGroupMap[i] = taskGroup.MapToTaskGroup()
  }

  return taskGroupMap, nil
}

