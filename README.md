# TaskTrackerCLI
The purpose of this project is to build a task tracker that runs using CLI commands

## How to run

### Build
Start by running ```go build``` in base directory

This should produce a binary file named **tasktracker** which will be the entry point to running the CLI.

### Run
Add New Task: ```./tasktracker add "Name of Task Here"```

Update Existing Task's Name: ```./tasktracker update 1 "New Task Name Here"```

Delete Task: ```./tasktracker delete 1```

Update Task's Status: ```./tasktracker mark-in-progress 1```
Update Task's Status: ```./tasktracker mark-done 1```

List All Tasks: ```./tasktracker list```
List Completed Tasks: ```./tasktracker list done```
List Tasks In Progress: ```./tasktracker list in-progress```

### Tests
