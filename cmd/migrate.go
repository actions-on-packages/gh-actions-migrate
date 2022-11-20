package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	// options := models.MigrateReleaseApiRequest{}
	var migrateCommand = &cobra.Command{
		Use:   "migrate",
		Short: "migrate action release",
		RunE: func(cmd *cobra.Command, args []string) error {
			// if len(args) != 0 {
			// 	return fmt.Errorf(fmt.Sprintf("Invalid argument(s). Expected 0 received %d", len(args)))
			// }
			// options, inputError := collectMissingInputs(options)

			// if inputError != nil {
			// 	fmt.Println(inputError)
			// 	return nil
			// }

			// var current_repo ghRepo.Repository
			// if options.Repo != "" {
			// 	var err error
			// 	current_repo, err = internal.GetRepo(options.Repo)
			// 	if err != nil {
			// 		return nil
			// 	}
			// }

			// fmt.Println("\nCreating new release...")

			// var response, err = service.CreateRelease(current_repo, options)
			// if err != nil {
			// 	fmt.Printf("\n%s", err)
			// 	return err
			// }

			// printResponse(current_repo, options, *response)
			return nil
		},
	}

	// createCommand.Flags().StringVarP(&options.ChangeType, "changetype", "c", "", "change type of the release")
	// createCommand.Flags().StringVarP(&options.Repo, "repo", "r", "", "repo to create release for")
	// createCommand.Flags().StringVarP(&options.Title, "title", "t", "", "title of the release")
	// createCommand.Flags().StringVarP(&options.ReleaseNotes, "releasenotes", "n", "", "release notes of the release")
	// createCommand.Flags().StringVarP(&options.PreRelease, "prerelease", "p", "", "is this a pre-release?")

	return migrateCommand
}
