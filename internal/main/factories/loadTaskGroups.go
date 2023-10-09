package factories

import (
	"log"

	data_usecases "github.com/joaovds/taskergo/internal/data/usecases/task_group"
	"github.com/joaovds/taskergo/internal/infra/database"
	repositories "github.com/joaovds/taskergo/internal/infra/repositories/task_groups"
	"github.com/joaovds/taskergo/internal/presentation/contracts"
	controllers "github.com/joaovds/taskergo/internal/presentation/controllers/task_group"
)

func MakeLoadTaskGroups() contracts.Controller {
  db, err := database.Connect()
  if err != nil {
    log.Fatal(err)
  }

  loadTaskGroupsRepository := repositories.NewLoadTaskGroupsRepository(db)
  loadTaskGroupsService := data_usecases.DbLoadTaskGroups(loadTaskGroupsRepository)
  loadTaskGroupsController := controllers.LoadTaskGroupsController{ LoadTaskGroups: loadTaskGroupsService }

  return loadTaskGroupsController
}

