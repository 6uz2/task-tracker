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

## Files

Tasks will be stored in JSON file under the same diectory, and the file name is `tasks.json`.
