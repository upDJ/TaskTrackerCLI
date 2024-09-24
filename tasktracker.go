package main

import (
	"fmt"
	"os"
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
	default:
		fmt.Println(
			"Usage Below \n ",
			"add -> to create a new task",
		)
	}

	task_manager.DumpTasksJson()
}
