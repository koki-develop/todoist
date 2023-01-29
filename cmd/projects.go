package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage projects",
	Long:  "Manage projects.",
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  "List all projects.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig(cmd)
		if err != nil {
			return ErrLoadConfig
		}

		cl := todoistapi.New(&todoistapi.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(*flagFormat.Get(cmd, false)))

		projs, err := cl.ListProjects()
		if err != nil {
			return err
		}

		o, err := rdr.Render(projs)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var projectsGetCmd = &cobra.Command{
	Use:   "get <PROJECT_ID>",
	Short: "Get a project",
	Long:  "Get a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig(cmd)
		if err != nil {
			return ErrLoadConfig
		}

		cl := todoistapi.New(&todoistapi.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(*flagFormat.Get(cmd, false)))

		proj, err := cl.GetProject(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(proj)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var projectsCreateCmd = &cobra.Command{
	Use:   "create <PROJECT_NAME>",
	Short: "Create a project",
	Long:  "Create a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := loadConfig(cmd)
		if err != nil {
			return ErrLoadConfig
		}

		cl := todoistapi.New(&todoistapi.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(*flagFormat.Get(cmd, false)))

		p := &todoistapi.CreateProjectPayload{
			Name:       name,
			ParentID:   flagProjectParentID.Get(cmd, true),
			IsFavorite: flagProjectFavorite.Get(cmd, true),
			Color:      flagProjectColor.Get(cmd, true),
		}

		proj, err := cl.CreateProject(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(proj)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var projectsUpdateCmd = &cobra.Command{
	Use:   "update <PROJECT_ID>",
	Short: "Update a project",
	Long:  "Update a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig(cmd)
		if err != nil {
			return ErrLoadConfig
		}

		cl := todoistapi.New(&todoistapi.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(*flagFormat.Get(cmd, false)))

		p := &todoistapi.UpdateProjectPayload{
			Name:       flagProjectName.Get(cmd, true),
			Color:      flagProjectColor.Get(cmd, true),
			IsFavorite: flagProjectFavorite.Get(cmd, true),
		}

		proj, err := cl.UpdateProject(id, p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(proj)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var projectsDeleteCmd = &cobra.Command{
	Use:   "delete <PROJECT_ID>",
	Short: "Delete a project",
	Long:  "Delete a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig(cmd)
		if err != nil {
			return ErrLoadConfig
		}

		cl := todoistapi.New(&todoistapi.Config{Token: cfg.APIToken})

		if err := cl.DeleteProject(id); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
		projectsDeleteCmd,
	)
}
