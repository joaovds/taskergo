package main

import "github.com/joaovds/taskergo/internal/cli"

func main() {
  cliFlags := cli.SetupCliFlags()
  cli.HandleCliFlags(cliFlags)
}

