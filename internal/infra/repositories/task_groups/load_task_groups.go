package repositories

import (
	"database/sql"

	repositories "github.com/joaovds/taskergo/internal/data/contracts/repositories/task_group"
	"github.com/joaovds/taskergo/internal/infra/queries"
)

type LoadTaskGroupsRepository struct {
  Db *sql.DB
}

func NewLoadTaskGroupsRepository(db *sql.DB) LoadTaskGroupsRepository {
  return LoadTaskGroupsRepository{Db: db}
}

func (l LoadTaskGroupsRepository) LoadAll() ([]repositories.LoadTaskGroupRepositoryResult, error) {
  var taskGroups []repositories.LoadTaskGroupRepositoryResult

  rows, err := l.Db.Query(queries.TaskGroups().GetTaskGroups)
  if err != nil {
    return taskGroups, err
  }
  defer rows.Close()

  for rows.Next() {
    var taskGroup repositories.LoadTaskGroupRepositoryResult

    err := rows.Scan(&taskGroup.Id, &taskGroup.Name, &taskGroup.Description)
    if err != nil {
      return taskGroups, err
    }

    taskGroups = append(taskGroups, taskGroup)
  }

  return taskGroups, nil
}

