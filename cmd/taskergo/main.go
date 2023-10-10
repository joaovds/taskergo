package main

import (
  "github.com/joaovds/taskergo/internal/main/cli"
  "github.com/joaovds/taskergo/internal/main/config"
)

func main() {
  cliFlags := config.SetupCliFlags()
  cli.HandleCliFlags(cliFlags)
}

