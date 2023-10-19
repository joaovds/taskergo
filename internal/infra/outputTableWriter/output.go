package outputtablewriter

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func WriteOutputTable(data [][]string, headers []string) {
  table := tablewriter.NewWriter(os.Stdout)

  table.SetHeader(headers)
  table.SetAutoFormatHeaders(true)

  table.AppendBulk(data)

  table.Render()
}

