package cmd

import (
	"fmt"
	"os"

	"github.com/Quasiterra/tinya/core"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "list",
	Short: "List all S3 buckets",
	// Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		if res, err := core.ListBuckets(); err == nil {
			fmt.Println(res)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
