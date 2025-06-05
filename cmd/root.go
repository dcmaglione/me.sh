package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "micro",
	Short: "Local encrypted microblogging",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
