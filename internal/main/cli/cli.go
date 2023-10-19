package cli

import (
  "flag"
  "fmt"

  "github.com/joaovds/taskergo/internal/main/config"
  "github.com/joaovds/taskergo/internal/main/factories"
)

func HandleCliFlags(cliFlags config.CliFlags) {
  if (*cliFlags.Help || (cliFlags.TotalFlags <= 0)) {
    flag.PrintDefaults()
  }

  if *cliFlags.ShowTaskGroups {
    fmt.Println(factories.MakeLoadTaskGroups().Exec())
  }
}

