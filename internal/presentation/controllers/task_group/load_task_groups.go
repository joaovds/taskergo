package controllers

import (
  usecases "github.com/joaovds/taskergo/internal/domain/usecases/task_group"
  "github.com/joaovds/taskergo/internal/presentation/contracts"
)

type LoadTaskGroupsController struct {
  loadTaskGroups usecases.LoadTaskGroups
}

func (c *LoadTaskGroupsController) Exec() contracts.HttpResponse {
  taskGroups, err := c.loadTaskGroups.Load()
  if err != nil {
    return contracts.ServerError(err.Error())
  }

  return contracts.Ok(taskGroups)
}

