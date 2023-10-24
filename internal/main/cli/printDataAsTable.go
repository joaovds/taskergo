package cli

import (
  "fmt"

  "github.com/joaovds/taskergo/internal/domain/entities"
  outputtablewriter "github.com/joaovds/taskergo/internal/infra/outputTableWriter"
)

func PrintDataAsTable(data interface{}) {
  outputData := make([][]string, 0)
  headers := make([]string, 0)

  switch values := data.(type) {
  case []entities.TaskGroup: 
    outputData, headers = handlePrintTaskGroups(values)

  default:
    panic(fmt.Sprintf("Unknown type %T", values))
}

  outputtablewriter.WriteOutputTable(outputData, headers)
}

func handlePrintTaskGroups(data []entities.TaskGroup) ([][]string, []string) {
  const uniqueIdentifier int = iota
  var outputData [][]string
  var headers []string

  for _, taskGroup := range data {
    outputData = append(outputData, []string{fmt.Sprintf("%d", uniqueIdentifier), taskGroup.Name, taskGroup.Description})
  }
  headers = entities.TaskGroup{}.GetTableHeaders()

  return outputData, headers
}

