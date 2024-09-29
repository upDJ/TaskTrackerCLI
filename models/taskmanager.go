package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// task manager -> list of tasks
type TaskManager struct {
	TaskList []Task
}

// read task from json
func (tm *TaskManager) ReadTaskJSON() TaskManager {
	var task_manager TaskManager
	var tasks []Task

	taskFile, err := os.Open("tasks.json")

	if err != nil {
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
func (tm *TaskManager) DumpTasksJson() {
	tasksJson, _ := json.Marshal(tm.TaskList)
	os.WriteFile("tasks.json", tasksJson, 0777)
}

// add task to the task manager
func (tm *TaskManager) AddTask(task_name string) {
	for _, task := range tm.TaskList {
		if task.Name == task_name {
			fmt.Printf("Task '%s' already exists. \n", task_name)
			return
		}
	}

	task := Task{ID: len(tm.TaskList) + 1, Name: task_name, Status: ""}
	tm.TaskList = append(tm.TaskList, task)
	fmt.Printf("Task added successfully (ID: %d)", task.ID)
}

// update task name
func (tm *TaskManager) UpdateTaskName(task_id int, new_task_name string) {
	for _, task := range tm.TaskList {
		if task.Name == new_task_name {
			fmt.Printf("Task '%s' already exists. \n", new_task_name)
			return
		}
	}

	for i, task := range tm.TaskList {
		if task.ID == task_id {
			tm.TaskList[i].Name = new_task_name
		}
	}
}

// delete task
func (tm *TaskManager) DeleteTask(task_id int) {
	for i, task := range tm.TaskList {
		if task.ID == task_id {
			fmt.Printf("Deleting task: %v", task.Name)
			tm.TaskList = append(tm.TaskList[:i], tm.TaskList[i+1:]...)
		}
	}
}

// update task status
func (tm *TaskManager) UpdateStatus(task_id int, status string) {
	for i, task := range tm.TaskList {
		if task.ID == task_id {
			tm.TaskList[i].Status = status
		}
	}
}


// list all
// list tasks that are done
// list tasks that are not done
// list tasks in progress

//modularize the block of code that checks for duplicates
