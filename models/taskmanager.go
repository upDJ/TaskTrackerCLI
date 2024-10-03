package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
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

	task := Task{ID: len(tm.TaskList) + 1, Name: task_name, Status: "", CreatedAt: time.Now(), UpdatedAt: time.Now()}
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
			tm.TaskList[i].UpdatedAt = time.Now()
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

// list all tasks
func (tm *TaskManager) ListTasks(status string) {
	fmt.Println("Tasks in Database")
	fmt.Println("-------------------")
	fmt.Println("Index\tName\t\tStatus\t\tCreated At\t\tUpdated At")

	if status == "" {
		for i, task := range tm.TaskList {
			fmt.Printf("%d.\t%s\t\t%s\t\t%s\t%s\n", i+1, task.Name, task.Status, task.CreatedAt.Format("2006-01-02T15:04:05"), task.UpdatedAt.Format("2006-01-02T15:04:05"))
		}
	} else {
		for i, task := range tm.TaskList {
			if task.Status == status {
				fmt.Printf("%d.\t%s\t\t%s\t\t%s\t%s\n", i+1, task.Name, task.Status, task.CreatedAt.Format("2006-01-02T15:04:05"), task.UpdatedAt.Format("2006-01-02T15:04:05"))
			}
		}
	}
}
