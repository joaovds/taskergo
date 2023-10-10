package config

import (
  "flag"
)

type CliFlags struct {
  Help *bool
  ShowTaskGroups *bool
}

func SetupCliFlags() CliFlags {
  cliFlags := CliFlags{
    Help: flag.Bool("help", false, "Show help"),
    ShowTaskGroups: flag.Bool("stg", false, "Show all Task groups"),
  }

  flag.Parse()

  return cliFlags
}

