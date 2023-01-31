package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string
)

var rootCmd = &cobra.Command{
	Use:  "todoist-cli",
	Long: "CLI Client for Todoist.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		switch err {
		case ErrLoadConfig, ErrAPITokenNotConfigured:
			fmt.Fprintln(os.Stderr, "Run `todoist-cli configure` to reconfigure.")
		}
		os.Exit(1)
	}
}

func init() {
	/*
	 * version
	 */
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version
	_ = notifyNewRelease(os.Stderr)

	/*
	 * add commands
	 */

	// configure
	rootCmd.AddCommand(configureCmd)

	// projects
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
		projectsDeleteCmd,
	)

	// sections
	rootCmd.AddCommand(sectionsCmd)
	sectionsCmd.AddCommand(
		sectionsListCmd,
		sectionsGetCmd,
		sectionsCreateCmd,
		sectionsUpdateCmd,
		sectionsDeleteCmd,
	)

	// tasks
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.AddCommand(
		tasksListCmd,
		tasksGetCmd,
		tasksCreateCmd,
	)

	/*
	 * basic flags
	 */

	// --api-token
	flagAPIToken.Add(configureCmd)
	flagAPIToken.Add(projectsCmd.Commands()...)
	flagAPIToken.Add(sectionsCmd.Commands()...)
	flagAPIToken.Add(tasksCmd.Commands()...)

	// -f, --format
	flagFormat.Add(configureCmd)
	flagFormat.Add(projectsCmd.Commands()...)
	flagFormat.Add(sectionsCmd.Commands()...)
	flagFormat.Add(tasksCmd.Commands()...)

	/*
	 * project flags
	 */

	// create
	flagProjectsCreateParentID.Add(projectsCreateCmd)
	flagProjectsCreateColor.Add(projectsCreateCmd)
	flagProjectsCreateFavorite.Add(projectsCreateCmd)

	// update
	flagProjectsUpdateColor.Add(projectsUpdateCmd)
	flagProjectsUpdateFavorite.Add(projectsUpdateCmd)
	flagProjectsUpdateName.Add(projectsUpdateCmd)

	/*
	 * section flags
	 */

	// list
	flagSectionsListProjectID.Add(sectionsListCmd)

	// create
	flagSectionsCreateProjectID.Add(sectionsCreateCmd)
	flagSectionsCreateOrder.Add(sectionsCreateCmd)

	// update
	flagSectionsUpdateName.Add(sectionsUpdateCmd)

	/*
	 * task flags
	 */

	// list
	flagTasksListProjectID.Add(tasksListCmd)
	flagTasksListSectionID.Add(tasksListCmd)
	flagTasksListLabel.Add(tasksListCmd)
	flagTasksListFilter.Add(tasksListCmd)
	flagTasksListLang.Add(tasksListCmd)
	flagTasksListIDs.Add(tasksListCmd)

	// create
	flagTasksCreateDescription.Add(tasksCreateCmd)
	flagTasksCreateProjectID.Add(tasksCreateCmd)
	flagTasksCreateSectionID.Add(tasksCreateCmd)
	flagTasksCreateParentID.Add(tasksCreateCmd)
	flagTasksCreateOrder.Add(tasksCreateCmd)
	flagTasksCreateLabels.Add(tasksCreateCmd)
	flagTasksCreatePriority.Add(tasksCreateCmd)
	flagTasksCreateDueString.Add(tasksCreateCmd)
	flagTasksCreateDueDate.Add(tasksCreateCmd)
	flagTasksCreateDueDatetime.Add(tasksCreateCmd)
	flagTasksCreateDueLang.Add(tasksCreateCmd)
	flagTasksCreateAssigneeID.Add(tasksCreateCmd)
}
