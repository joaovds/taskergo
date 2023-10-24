package queries

type TaskGroupsQueries struct {
  GetTaskGroups string
}

func TaskGroups() TaskGroupsQueries {
  TaskGroups := TaskGroupsQueries{
    GetTaskGroups: `SELECT id, name, description FROM task_groups`,
  }

  return TaskGroups
}

