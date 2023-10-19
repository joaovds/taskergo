package config

import (
  "flag"
)

type CliFlags struct {
  TotalFlags int
  Help *bool
  ShowTaskGroups *bool
}

func SetupCliFlags() CliFlags {
  cliFlags := CliFlags{
    Help: flag.Bool("help", false, "Show help"),
    ShowTaskGroups: flag.Bool("stg", false, "Show all Task groups"),
  }

  flag.Parse()

  cliFlags.TotalFlags = flag.NFlag()

  return cliFlags
}

