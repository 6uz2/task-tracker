package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var taskFilePath string = "tasks.json"

type TaskCollection struct {
	LastTaskId int
	Tasks      []TaskProperties
}

type TaskProperties struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: task-cli [command] [task ID] [description | option]\n")
	flag.PrintDefaults()
}

func readTasksFile() *TaskCollection {
	var existingTasks TaskCollection
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

func writeTasksFile(tasks *TaskCollection) {
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

// Remove element at index from the given slice
func removeTaskProperty(slice []TaskProperties, s int) []TaskProperties {
	return append(slice[:s], slice[s+1:]...)
}

func getTaskIndex(tasks []TaskProperties, taskId int) int {
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

func addTask() *cobra.Command {
	var addTask = &cobra.Command{
		Use:   "add",
		Short: "Add a new task and return the task ID.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			currentTime := time.Now().UTC()
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli add \"task description\"")
				return
			}

			taskDescription := cmdArgs[1]
			taskData := readTasksFile()

			newTask := TaskProperties{
				Id:          taskData.LastTaskId + 1,
				Description: taskDescription,
				Status:      "todo",
				CreatedAt:   currentTime,
				UpdatedAt:   currentTime,
			}

			taskData.Tasks = append(taskData.Tasks, newTask)

			taskData.LastTaskId = newTask.Id

			writeTasksFile(taskData)

			fmt.Printf("Task added successfully (ID: %d)\n", newTask.Id)
		},
	}

	return addTask
}

// Update task information
func updateTask() *cobra.Command {
	var updateTask = &cobra.Command{
		Use:   "update",
		Short: "Update the task info",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdArgs := flag.Args()

			if len(cmdArgs) != 3 {
				fmt.Println("Usage: task-cli update [task Id] \"task description\"")
				return
			}

			updateTaskId, err := strconv.Atoi(cmdArgs[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			newDescription := cmdArgs[2]

			taskData := readTasksFile()

			if len(taskData.Tasks) == 0 {
				fmt.Println("Currently no tasks.")
				return
			}

			if taskIdx := getTaskIndex(taskData.Tasks, updateTaskId); taskIdx < 0 {
				return
			} else {
				taskData.Tasks[taskIdx].Description = newDescription
				taskData.Tasks[taskIdx].UpdatedAt = time.Now().UTC()
			}

			writeTasksFile(taskData)

			fmt.Printf("Task is updated successfully (ID: %d)\n", updateTaskId)
		},
	}

	return updateTask
}

// Delete task command
func deleteTask() *cobra.Command {
	var deleteTask = &cobra.Command{
		Use:   "delete",
		Short: "Delete task by the given task ID.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli delete [task Id]")
				return
			}

			deleteTaskId, err := strconv.Atoi(cmdArgs[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			taskData := readTasksFile()

			if len(taskData.Tasks) == 0 {
				fmt.Println("Currently no tasks.")
				return
			}

			if taskIdx := getTaskIndex(taskData.Tasks, deleteTaskId); taskIdx < 0 {
				return
			} else {
				taskData.Tasks = removeTaskProperty(taskData.Tasks, taskIdx)
			}

			writeTasksFile(taskData)

			fmt.Printf("Task is deleted successfully (ID: %d)\n", deleteTaskId)
		},
	}

	return deleteTask
}

// List task by status
func listTasks() *cobra.Command {
	var listTasks = &cobra.Command{
		Use:   "list",
		Short: "Listing task by status.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return listTasks
}

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "task-cli is a tool to track and manage your task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmdArgs := flag.Args()

		if len(cmdArgs) <= 1 {
			flag.Usage()
			return
		}
	},
}

func Execute() error {
	rootCmd.AddCommand(addTask())
	rootCmd.AddCommand(updateTask())
	rootCmd.AddCommand(deleteTask())
	rootCmd.AddCommand(listTasks())
	return rootCmd.Execute()
}
