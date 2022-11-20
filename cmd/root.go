package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "action-migrate",
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func Execute() {
	addSubCommandsToRoot()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandsToRoot() {
	rootCmd.AddCommand(NewCmdMigrate())
}

func getRootHelp() string {
	return `
gh-actions-migrate: Works with GitHub Actions releases. 
USAGE:
	gh-action-migrate <command> [flags]
CORE COMMANDS:
	create:		migrates action release

FLAGS
	--help                      Show help for create command
EXAMPLES:
	$ gh action-release create
	$ gh action-release create -r github/js-action -c minor -t release title -p n
	$ gh action-release create -r github/js-action -c minor -t release title -n release notes -p n
	$ gh action-release create --repo github/js-action --changetype minor --title release title --releasenotes release notes --prerelease n
`
}
