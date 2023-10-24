package cli

import (
  "flag"

  "github.com/joaovds/taskergo/internal/main/config"
  "github.com/joaovds/taskergo/internal/main/factories"
)

func HandleCliFlags(cliFlags config.CliFlags) {
  if (*cliFlags.Help || (cliFlags.TotalFlags <= 0)) {
    flag.PrintDefaults()
  }

  if *cliFlags.ShowTaskGroups {
    taskGroups := factories.MakeLoadTaskGroups().Exec()

    PrintDataAsTable(taskGroups.Data)
  }
}

