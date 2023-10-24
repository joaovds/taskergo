package entities

type TaskGroup struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
}

func (taskGroup TaskGroup) GetTableHeaders() []string {
  return []string{"", "Name", "Description"}
}

