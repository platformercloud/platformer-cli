package set

import (
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

// SetCmd is the base command for all resource set commands
// set is the same as 'selectprompt' but without the prompt list.
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set your default organization and project",
	ValidArgs: []string{
		organizationSetCmd.Use,
		projectSetCmd.Use,
	},
	ArgAliases: append(
		organizationSetCmd.Aliases,
		projectSetCmd.Aliases...,
	),
	Args: cobra.ExactValidArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Runs before all child commands (eg. project/org list)
		cli.HandleErrorAndExit(func() error {
			if !auth.IsLoggedIn() {
				return &cli.NotLoggedInError{}
			}
			return nil
		}())
	},
}

func init() {
	SetCmd.AddCommand(organizationSetCmd)
	SetCmd.AddCommand(projectSetCmd)
}
