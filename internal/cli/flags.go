package cli

import "flag"

type CliFlags struct {
  help *bool
}

func SetupCliFlags() CliFlags {
  cliFlags := CliFlags{
    help: flag.Bool("help", false, "Show help"),
  }

  flag.Parse()

  return cliFlags
}

func HandleCliFlags(cliFlags CliFlags) {
  if *cliFlags.help {
    flag.PrintDefaults()
  }
}

