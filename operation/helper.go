package operation

import (
	"errors"
	"fmt"
	"task-tracker/task-tracker/dto"
	"time"
)

func getTaskIndex(tasks []dto.TaskProperties, taskId int) int {
	var foundTaskIdx = -1
	for idx, task := range tasks {
		if task.Id == taskId {
			foundTaskIdx = idx
			break
		}
	}

	if foundTaskIdx < 0 {
		fmt.Println("No task found.")
	}

	return foundTaskIdx
}

// Remove element at index from the given slice
func removeTaskProperty(slice []dto.TaskProperties, s int) []dto.TaskProperties {
	return append(slice[:s], slice[s+1:]...)
}

func CreateTask(desc string) (taskId int, err error) {
	taskData := readTasksFile()

	if taskData == nil {
		return -1, errors.New("failed to read task file")
	}

	currentTime := time.Now().UTC()

	newTask := dto.TaskProperties{

		Id:          taskData.LastTaskId + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	taskData.Tasks = append(taskData.Tasks, newTask)

	taskData.LastTaskId = newTask.Id

	writeTasksFile(taskData)

	fmt.Printf("Task added successfully (ID: %d)\n", taskId)

	return newTask.Id, nil
}

func UpdateTaskDescription(taskId int, newDescription string) error {
	taskData := readTasksFile()

	if taskData == nil {
		return errors.New("failed to read task file")
	}

	if len(taskData.Tasks) == 0 {
		fmt.Println("Currently no tasks.")
		return nil
	}

	if taskIdx := getTaskIndex(taskData.Tasks, taskId); taskIdx < 0 {
		return nil
	} else {
		taskData.Tasks[taskIdx].Description = newDescription
		taskData.Tasks[taskIdx].UpdatedAt = time.Now().UTC()
	}

	writeTasksFile(taskData)

	fmt.Printf("Task description is updated (ID: %d)\n", taskId)

	return nil
}

func DeleteTask(taskId int) error {
	taskData := readTasksFile()

	if taskData == nil {
		return errors.New("failed to read task file")
	}

	if len(taskData.Tasks) == 0 {
		fmt.Println("Currently no tasks.")
		return nil
	}

	if taskIdx := getTaskIndex(taskData.Tasks, taskId); taskIdx < 0 {
		return nil
	} else {
		taskData.Tasks = removeTaskProperty(taskData.Tasks, taskIdx)
	}

	writeTasksFile(taskData)

	fmt.Printf("Task is deleted successfully (ID: %d)\n", taskId)

	return nil
}
