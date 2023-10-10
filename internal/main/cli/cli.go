package cli

import (
  "flag"
  "fmt"

  "github.com/joaovds/taskergo/internal/main/config"
  "github.com/joaovds/taskergo/internal/main/factories"
)

func HandleCliFlags(cliFlags config.CliFlags) {
  if *cliFlags.Help {
    flag.PrintDefaults()
  }

  a := factories.MakeLoadTaskGroups().Exec()
  factories.MakeLoadTaskGroups().Exec()
  factories.MakeLoadTaskGroups().Exec()
  factories.MakeLoadTaskGroups().Exec()
  fmt.Println(a)
}

