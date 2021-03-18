package cmd

import (
        "fmt"
        "github.com/spf13/cobra"
        "os"
)

var manualEntryCmd = &cobra.Command{
        Use: "",
}


var rootCmd = &cobra.Command{
        Use: "timetracker-cli",
}

func Execute() {
        rootCmd.AddCommand(manualEntryCmd)

        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

