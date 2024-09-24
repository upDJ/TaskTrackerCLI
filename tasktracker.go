package main

import (
	"io"
	"os"
    "fmt"
    "flag"
    "encoding/json"
	"tasktracker/models"
)

func main() {
    taskFile, err := os.Open("tasks.json")
    

    var taskManager models.TaskManager
    if err != nil{
        taskManager = models.TaskManager{TaskList: make([]models.Task, 3)}
    } else {
        taskData, err := io.ReadAll(taskFile)
        if err != nil {
            fmt.Printf("Unable to read tasks.json, error: %v", err)
            return 
        }
    
        var tasks []models.Task
        json.Unmarshal([]byte(taskData), &tasks)
        taskManager = models.TaskManager{TaskList: tasks}
    }

    defer taskFile.Close()


    task := flag.String("a", "", "Enter New Task Name")
    flag.Parse()

    if (*task != "") {
        taskManager.AddTask(*task)
    }

    taskManager.DumpTasksJson()
}