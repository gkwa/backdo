package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/backdo/test1"
)

var (
	dirPath1        string
	dirPath2        string
	excludeExisting []string
	script          bool
)

// test1Cmd represents the test1 command
var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dirPath1 == "" || dirPath2 == "" {
			return fmt.Errorf("both directory paths (--incoming and --existing) are required")
		}
		if script {
			err := test1.GenerateScript(dirPath1, dirPath2, excludeExisting)
			if err != nil {
				return err
			}
		} else {
			err := test1.RunTest(dirPath1, dirPath2, excludeExisting)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)

	test1Cmd.Flags().StringVar(&dirPath1, "incoming", "", "First directory path")
	test1Cmd.Flags().StringVar(&dirPath2, "existing", "", "Second directory path")
	test1Cmd.Flags().StringSliceVar(&excludeExisting, "exclude-existing", nil, "List of substrings to exclude from existing directory paths")
	test1Cmd.Flags().BoolVar(&script, "script", false, "Generate bash CLI commands")
}
