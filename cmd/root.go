package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "p2p file transfer",
	Short: "Ipfs CLI p2p file transfer made easy",
	Long:  "This CLI allows you to transfer your local file to your peers via IPFS",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error executing: ", err)
		os.Exit(1)
	}
}
