package cmd

import (
	"fmt"
	"os"

	"github.com/naftulikay/gobender/etl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobender",
	Short: "A simple ETL worker demo in Go.",
	Long:  "Really, just a simple ETL worker demo in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		etl.Run()
	},
}

// Execute is the main entrypoint to the application into Cobra.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
