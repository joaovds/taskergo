package outputtablewriter

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func WriteOutputTable(data [][]string, headers []string) {
  table := tablewriter.NewWriter(os.Stdout)

  table.SetHeader(headers)
  table.SetAutoFormatHeaders(true)

  fmt.Printf("%T", data)

  table.AppendBulk(data)

  table.Render()
}

