package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	api "github.com/ipfs/go-ipfs-api"
	"github.com/moon004/p2p-sharer/cnf"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "p2p-sharer",
		Short: "Simplify the file transfer protocol to other machine via IPFS protocol",
		Long: `
This CLI allows you to transfer your local file to your peers via IPFS protocol
Also it enables you to add your peers to your friend list and make consequent
file transfer more easy and faster.

IMPORTANT!!

Make sure you host your own node by running before using any of the command:

	'ipfs daemon'`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			startTime = time.Now()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				timeSpent := int64(time.Since(startTime) / time.Millisecond)
				d.Info(fmt.Sprintf("Time Spent: %vms", timeSpent))
			}
		},
	}

	startTime time.Time
)

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
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "output verbose")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "trigger debuging mode")
	//									Duration takes int64
	rootCmd.PersistentFlags().Int64("timeout", cnf.Timeout, "timeout")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.AddCommand(
		UpFile(),
		GetObject(),
		FriendList(),
		AddFriend(),
		MyIdentity(),
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
	viper.ReadInConfig()
}

// NewIpfsAPI returns ipfs api shell
func NewIpfsAPI() *api.Shell {
	return api.NewLocalShell()
}
