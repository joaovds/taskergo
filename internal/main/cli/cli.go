package cli

import (
	"flag"

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

    // taskGroups.Data
    // preciso deixar esse taskGroups.Data como [][]string
    normalizedTaskGroups := [][]string{
    }

    outputtablewriter.WriteOutputTable(taskGroups.Data, []string{"oie", "oie2", "oie3"})
  }
}

