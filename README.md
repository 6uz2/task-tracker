# Task Tracker

Solution for [task-tracker](https://roadmap.sh/projects/task-tracker) project from [roadmap.sh](https://roadmap.sh/).

## Installation

```shell
git clone https://github.com/6uz2/task-tracker.git; cd task-tracker;

go mod vendor

go build -o task-cli main.go

# See usage below and run the task-cli
```

## Usage

### Add Task

```shell
# Adding a new task
./task-cli add "task description"
# Output: Task added successfully (ID: 1)
```

### Update Task Description

```shell
./task-cli update <Task ID> "Buy groceries and cook dinner"
```

### Update Task Status

```shell
# Marking a task as in progress or done
./task-cli mark-in-progress <Task ID>
./task-cli mark-done <Tasks ID>
```

### Delete Task

```shell
# Delete Task ID: 1
./task-cli delete <Tasks ID>
```

### List Tasks

```shell
# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

#### Examples

```shell
# Example of listing all tasks

# Create tasks
./task-cli add "Task #1"
./task-cli add "Task #2"
./task-cli add "Task #3"

# List all tasks
./task-cli list
# Output:
# [
#   {
#     "Id": 1,
#     "Description": "Task #1",
#     "Status": "todo",
#     "CreatedAt": "2025-01-17T11:22:50.15535Z",
#     "UpdatedAt": "2025-01-17T11:22:50.15535Z"
#   },
#   {
#     "Id": 2,
#     "Description": "Task #2",
#     "Status": "todo",
#     "CreatedAt": "2025-01-17T11:22:51.900836Z",
#     "UpdatedAt": "2025-01-17T11:22:51.900836Z"
#   },
#   {
#     "Id": 3,
#     "Description": "Task #3",
#     "Status": "todo",
#     "CreatedAt": "2025-01-17T11:22:53.40701Z",
#     "UpdatedAt": "2025-01-17T11:22:53.40701Z"
#   }
# ]
```

```shell
# Example of listing all tasks

# Create tasks
./task-cli add "Task #1"
./task-cli add "Task #2"
./task-cli add "Task #3"

# Mark Task #1 to "done" and Task #2 to "in-progress"
./task-cli mark-in-progress 2
./task-cli mark-done 1

# List "done" tasks
./task-cli list "done"
# Output: 
# [
#   {
#     "Id": 1,
#     "Description": "Task #1",
#     "Status": "done",
#     "CreatedAt": "2025-01-17T11:22:50.15535Z",
#     "UpdatedAt": "2025-01-17T11:24:57.990762Z"
#   }
# ]

# List "in-progress" tasks
# Output:
# [
#   {
#     "Id": 2,
#     "Description": "Task #2",
#     "Status": "in-progress",
#     "CreatedAt": "2025-01-17T11:22:51.900836Z",
#     "UpdatedAt": "2025-01-17T11:24:50.712536Z"
#   }
# ]
```

## Files

Tasks will be stored in JSON file under the same diectory, the file name is `tasks.json`.

