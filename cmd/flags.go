package cmd

import (
	"fmt"
	"strings"

	"github.com/koki-develop/todoist/pkg/renderer"
	"github.com/spf13/cobra"
)

var (
	apiToken string
	format   string
)

func init() {
	// api token
	for _, cmd := range []*cobra.Command{
		configureCmd,
		// projects
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
	} {
		cmd.Flags().StringVar(&apiToken, "api-token", "", "todoist api token")
	}

	// format
	for _, cmd := range []*cobra.Command{
		// projects
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
	} {
		cmd.Flags().StringVarP(&format, "format", "f", string(renderer.FormatTable), fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|")))
	}
}
