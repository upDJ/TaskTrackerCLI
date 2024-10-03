package models

import (
	"testing"
)

// func Initialize_Task_Manager(t *testing.T) {
// 	t.Helper()

// 	tm := TaskManager{}

// }

func TestAddTaskLength(t *testing.T) {
	tm := TaskManager{}
	tm.AddTask("Test Task List Length")
	task_list := tm.TaskList
	
	if len(task_list) != 1{
		t.Errorf("Length of Task List was incorrect, got: %d, want: %d", len(task_list), 1)
	}
}

func TestAddTask(t *testing.T) {
	tm := TaskManager{}
	tm.AddTask("Test Add Task")
	
	test_task := tm.TaskList[0]
	sample_task := Task{ID: 1, Name: "Test Add Task", Status: "", CreatedAt: test_task.CreatedAt, UpdatedAt: test_task.UpdatedAt}
	
	if sample_task != test_task{
		t.Errorf("Sample Task and Test Task are not equal, Sample Task: %+v, Added Task: %+v", sample_task, test_task)
	}

}

func TestUpdateTaskName(t *testing.T) {
	tm := TaskManager{}
	tm.AddTask("Test Update Task")

	updated_task_name := "Updated Task Name"
	tm.UpdateTaskName(1, updated_task_name)
	
	updated_task := tm.TaskList[0]
	
	if updated_task.Name != updated_task_name {
		t.Errorf("Task name has not been updated correctly, exepcted: %s, got: %s", updated_task.Name, updated_task_name)
	}
}


func TestUpdateStatus(t *testing.T) {
	tm := TaskManager{}
	tm.AddTask("Test Update Task Status")

	updated_task_status := "done"
	tm.UpdateStatus(1, updated_task_status)

	updated_task := tm.TaskList[0]
	
	if updated_task.Status != updated_task_status {
		t.Errorf("Task status has not been updated correctly, got: %s, expected: %s", updated_task.Status, updated_task_status)
	}

}

func TestDeleteTask(t *testing.T) {
	tm := TaskManager{}
	tm.AddTask("Test Delete Task")

	tm.DeleteTask(1)
	if len(tm.TaskList) != 0 {
		t.Errorf("Task has not been deleted, expected empty task list, got: %+v", tm.TaskList)
	}
}
