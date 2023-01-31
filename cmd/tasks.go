package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage tasks",
	Long:  "Manage tasks.",
}

var tasksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all active tasks",
	Long:  "List all active tasks.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListTasksParameters{
			ProjectID: flagTasksListProjectID.Get(cmd, true),
			SectionID: flagTasksListSectionID.Get(cmd, true),
			Label:     flagTasksListLabel.Get(cmd, true),
			Filter:    flagTasksListFilter.Get(cmd, true),
			Lang:      flagTasksListLang.Get(cmd, true),
			IDs:       flagTasksListIDs.Get(cmd, true),
		}
		ts, err := client.ListTasks(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(ts)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksGetCmd = &cobra.Command{
	Use:   "get <TASK_ID>",
	Short: "Get an active task",
	Long:  "Get an active task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		t, err := client.GetTask(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(t)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksCreateCmd = &cobra.Command{
	Use:   "create <TASK_CONTENT>",
	Short: "Create a task",
	Long:  "Create a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.CreateTaskParameters{
			Content:     &content,
			Description: flagTasksCreateDescription.Get(cmd, true),
			ProjectID:   flagTasksCreateProjectID.Get(cmd, true),
			SectionID:   flagTasksCreateSectionID.Get(cmd, true),
			ParentID:    flagTasksCreateParentID.Get(cmd, true),
			Order:       flagTasksCreateOrder.Get(cmd, true),
			Labels:      flagTasksCreateLabels.Get(cmd, true),
			Priority:    flagTasksCreatePriority.Get(cmd, true),
			DueString:   flagTasksCreateDueString.Get(cmd, true),
			DueDate:     flagTasksCreateDueDate.Get(cmd, true),
			DueDatetime: flagTasksCreateDueDatetime.Get(cmd, true),
			DueLang:     flagTasksCreateDueLang.Get(cmd, true),
			AssigneeID:  flagTasksCreateAssigneeID.Get(cmd, true),
		}
		t, err := client.CreateTask(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(t)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}
