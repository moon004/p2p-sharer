package cmd

import (
	"log"
	"os"
	"time"

	"github.com/moon004/p2p-sharer/cnf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "p2p-sharer",
	Short: "Simplify the file transfer protocol to other machine via IPFS protocol",
	Long: `
This CLI allows you to transfer your local file to your peers via IPFS protocol

Make sure you host your own node before using this CLI by invoking:

	ipfs daemon`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		startTime = time.Now()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			timeSpent := int64(time.Now().Sub(startTime) / time.Millisecond)
			log.Printf("%s", timeSpent)
		}
	},
}

var startTime time.Time
var Verbose bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().BoolP("verbose", "v", false, "output verbose")
	rootCmd.Flags().BoolP("debug", "d", false, "trigger debuging mode")

	rootCmd.AddCommand(
		SendFile(),
		AddFile(),
		ConnPeers(),
		FriendList(),
		AddFriend(),
	)
}

func initConfig() {
	configFile := cnf.ConfigStruct{}
	file := configFile.ConfigFile()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config found. Creating new config file")
		err = configFile.DefaultConfigValue()
	}
}
