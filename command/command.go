package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var taskFilePath string = "tasks.json"

type TaskCollection struct {
	Tasks []TaskProperties
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

func addTask() *cobra.Command {
	var addTask = &cobra.Command{
		Use:   "add",
		Short: "Add a new task and return the task ID.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			currentTime := time.Now().UTC()
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli add \"blablabla\"")
				return
			}

			taskDescription := cmdArgs[1]
			taskData := readTasksFile()

			newTask := TaskProperties{
				Id:          len(taskData.Tasks) + 1,
				Description: taskDescription,
				Status:      "todo",
				CreatedAt:   currentTime,
				UpdatedAt:   currentTime,
			}

			taskData.Tasks = append(taskData.Tasks, newTask)

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
