package cli

import (
	"flag"
	"fmt"

	"github.com/joaovds/taskergo/internal/domain/entities"
	outputtablewriter "github.com/joaovds/taskergo/internal/infra/outputTableWriter"
	"github.com/joaovds/taskergo/internal/main/config"
	"github.com/joaovds/taskergo/internal/main/factories"
)

func HandleCliFlags(cliFlags config.CliFlags) {
  if (*cliFlags.Help || (cliFlags.TotalFlags <= 0)) {
    flag.PrintDefaults()
  }

  if *cliFlags.ShowTaskGroups {
    taskGroups := factories.MakeLoadTaskGroups().Exec()

    var data []entities.TaskGroup
    if convertedTaskGroupData, ok := taskGroups.Data.([]entities.TaskGroup); ok {
      data = convertedTaskGroupData
    }

    var outputData [][]string
    for _, taskGroup := range data {
      outputData = append(outputData, []string{taskGroup.Name, taskGroup.Description})
    }
    headers := []string{"Name", "Description"}

    fmt.Println("Task Groups:")
    outputtablewriter.WriteOutputTable(outputData, headers)
  }
}

