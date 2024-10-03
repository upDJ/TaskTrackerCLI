package main

import (
	"fmt"
	"os"
	"strconv"
	"tasktracker/models"
)

func main() {
	var tm models.TaskManager

	task_manager := tm.ReadTaskJSON()
	args := os.Args

	switch args[1] {
	case "add":
		new_task := args[2]
		task_manager.AddTask(new_task)
	case "update":
		task_id, _ := strconv.Atoi(args[2])
		task_name := args[3]
		task_manager.UpdateTaskName(task_id, task_name)
	case "delete":
		task_id, _ := strconv.Atoi(args[2])
		task_manager.DeleteTask(task_id)
	case "mark-in-progress":
		task_id, _ := strconv.Atoi(args[2])
		status := "in progress"
		task_manager.UpdateStatus(task_id, status)
	case "mark-done":
		task_id, _ := strconv.Atoi(args[2])
		status := "done"
		task_manager.UpdateStatus(task_id, status)
	case "list":
		status := ""
		if len(args) > 2 {
			status = args[2]
		}
		task_manager.ListTasks(status)
	default:
		fmt.Println(
			"Usage Below \n ",
			"add -> to create a new task",
		)
	}

	task_manager.DumpTasksJson()
}
