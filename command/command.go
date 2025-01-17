package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	op "task-tracker/task-tracker/operation"

	"github.com/spf13/cobra"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: task-cli [command] [task ID] [description | option]\n")
	flag.PrintDefaults()
}

func addTask() *cobra.Command {
	var addTask = &cobra.Command{
		Use:   "add",
		Short: "Add a new task and return the task ID.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli add \"task description\"")
				return
			}

			taskDescription := cmdArgs[1]

			if _, err := op.CreateTask(taskDescription); err != nil {
				return
			}
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

			if err := op.UpdateTaskDescription(updateTaskId, newDescription); err != nil {
				return
			}
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

			if err := op.DeleteTask(deleteTaskId); err != nil {
				return
			}
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

func markTaskInProgress() *cobra.Command {
	var markTaskInProgress = &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark task is in progress status.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli mark-in-progress [task Id]")
				return
			}

			taskId, err := strconv.Atoi(cmdArgs[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if err := op.UpdateTaskStatus(taskId, "in-progress"); err != nil {
				return
			}
		},
	}

	return markTaskInProgress
}

func markTaskDone() *cobra.Command {
	var markTaskDone = &cobra.Command{
		Use:   "mark-done",
		Short: "Mark task is done.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdArgs := flag.Args()

			if len(cmdArgs) != 2 {
				fmt.Println("Usage: task-cli mark-in-progress [task Id]")
				return
			}

			taskId, err := strconv.Atoi(cmdArgs[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if err := op.UpdateTaskStatus(taskId, "done"); err != nil {
				return
			}
		},
	}

	return markTaskDone
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
	rootCmd.AddCommand(markTaskInProgress())
	rootCmd.AddCommand(markTaskDone())
	return rootCmd.Execute()
}
