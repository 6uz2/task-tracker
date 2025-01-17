package operation

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/task-tracker/dto"
)

var taskFilePath string = "tasks.json"

func readTasksFile() *dto.TaskCollection {
	var existingTasks dto.TaskCollection
	if _, err := os.Stat(taskFilePath); !os.IsNotExist(err) {
		fileBytes, err := os.ReadFile(taskFilePath)
		if err != nil {
			log.Fatalf("Error when opening %s file: %s", taskFilePath, err)
		}

		if len(fileBytes) > 0 {
			err = json.Unmarshal(fileBytes, &existingTasks)
			if err != nil {
				fmt.Println("Error unmarshalling data:", err)
				return nil
			}
		} else {
			fmt.Println("Empty file")
		}

	} else if os.IsNotExist(err) {
		file, err := os.Create(taskFilePath)
		if err != nil {
			log.Fatalf("Error creating %s file: %s", taskFilePath, err)
		}
		defer file.Close()
	}

	return &existingTasks
}

func writeTasksFile(tasks *dto.TaskCollection) {
	marshaled, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println("Error marshalling taskCollection: ", err)
		return
	}

	err = os.WriteFile(taskFilePath, marshaled, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}
