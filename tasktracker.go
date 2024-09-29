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
		task_number, _ := strconv.Atoi(args[2])
		task_name := args[3]
		task_manager.UpdateTaskName(task_number, task_name)
	case "delete":
		task_number, _ := strconv.Atoi(args[2])
		task_manager.DeleteTask(task_number)
	default:
		fmt.Println(
			"Usage Below \n ",
			"add -> to create a new task",
		)
	}

	task_manager.DumpTasksJson()
}
