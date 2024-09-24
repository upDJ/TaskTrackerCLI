package models

import (
	"os"
	"fmt"
	"encoding/json"
)

// task manager -> list of tasks
type TaskManager struct {
	TaskList []Task
}

// dump task to json
func (tm *TaskManager) DumpTasksJson(){
	tasksJson, _ := json.Marshal(tm.TaskList)
	os.WriteFile("tasks.json", tasksJson, 0777)
}

// add task to the task manager
func (tm *TaskManager) AddTask(taskname string) {
	for _, task := range tm.TaskList {
		if task.Name == taskname {
			fmt.Printf("Task '%s' already exists. \n", taskname)
			return 
		}
	}

	task := Task{Name: taskname, Status: ""}
	tm.TaskList = append(tm.TaskList, task)
}

// update task

// delete, manage array size
 
// mark a task as in progress or done
// list all
// list tasks that are done
// list tasks that are not done
// list tasks in progress

