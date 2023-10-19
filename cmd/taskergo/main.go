package main

import (
	"fmt"

	"github.com/joaovds/taskergo/internal/main/cli"
	"github.com/joaovds/taskergo/internal/main/config"
)

func main() {
  fmt.Println("TaskerGo - Welcome!")
  fmt.Println("")

  cliFlags := config.SetupCliFlags()
  cli.HandleCliFlags(cliFlags)
}

