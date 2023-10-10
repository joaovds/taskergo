package config

import (
  "flag"
)

type CliFlags struct {
  Help *bool
}

func SetupCliFlags() CliFlags {
  cliFlags := CliFlags{
    Help: flag.Bool("help", false, "Show help"),
  }

  flag.Parse()

  return cliFlags
}

