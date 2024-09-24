package models

import (
	"io"
	"os"
	"fmt"
	"encoding/json"
)

// task manager -> list of tasks
type TaskManager struct {
	TaskList []Task
}

// read task from json
func (tm *TaskManager) ReadTaskJSON() TaskManager{
	var task_manager TaskManager
	var tasks []Task

	taskFile, err := os.Open("tasks.json")

	if err != nil{
        task_manager = TaskManager{TaskList: []Task{}}
		return task_manager
    } 
	defer taskFile.Close()

	taskData, _ := io.ReadAll(taskFile)
	json.Unmarshal([]byte(taskData), &tasks)

	task_manager = TaskManager{TaskList: tasks}
	return task_manager
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

// update task status

// delete, manage array size
 
// mark a task as in progress or done
// list all
// list tasks that are done
// list tasks that are not done
// list tasks in progress

