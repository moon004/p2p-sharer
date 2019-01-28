// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/moon004/p2p-sharer/cnf"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetObject() *cobra.Command {
	var getobject = &cobra.Command{
		Use:   "retobject",
		Short: "Retrieve file from peers",
		Long: `Retrieve file from peers and save to local directory.
Make sure to add him/her as friend before retrieve the file.

Examples:
	
	` + tools.Args0() + ` retobject QmSgc9oPMqBppGyM3TWc7NZF11bwH8o3CDekd6pAYGJF8X "Siang Hwa" -o "Avengers.MP4"`,

		Args: cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			allflags := cmd.Flags()
			if allflags.Changed("fileName") == false {
				d.OnError(errors.New("Must provide value for all the required flag"))
				return
			}
			retobject(cmd, args)
		},
	}

	getobject.Flags().SortFlags = false
	getobject.Flags().StringP("fileName", "o", "", "The output filename (required)")
	return getobject
}

func retobject(cmd *cobra.Command, args []string) {
	friendName := args[1]
	fileName, _ := cmd.Flags().GetString("fileName")
	hash := args[0]
	regHash, _ := tools.Regex("ipfs hash")
	if !regHash.MatchString(hash) {
		d.OnError(errors.New("Invalid hash format"))
	}
	/*
		1. if got provide friendName, find friendName in config file
		2. if provided friendName but cant find == Not Provided == Dont Have
		3. if Dont Have, just dht findprovs, and connect and get <hash>
		4. If findprovs empty, just get <hash>
		5. if Have PeerID, Connect to it
	*/
	sh := NewIpfsAPI()
	// if -n is NOT empty
	if friendName != "" {
		// Acquire the friend's Peer's ID
		var c cnf.ConfigStruct
		Flist, err := c.GetFList()
		d.OnError(err)
		PeerInfo := Flist[friendName]

		if PeerInfo == "" {
			// Induce an error message
			err := errors.Errorf("You have no such friend: %s", friendName)
			d.OnError(err) // stop program here and show error
		}
		// Do step 5.																							// 2 minutes
		ctx, cancel := context.WithTimeout(context.Background(), tools.GetTimeout())
		defer cancel()
		err = sh.SwarmConnect(ctx, PeerInfo)
		d.OnError(err)
	}
	// Output to .p2p-sharer/storage
	p2pPath, _ := viper.Get("p2p_config_file").(string)
	dir, _ := filepath.Split(p2pPath) // dir == ../.p2p-sharer
	filePath := filepath.Join(dir, "storage")
	if _, err := os.Stat(filePath); err != nil && os.IsNotExist(err) {
		// If not exists create it
		os.Mkdir(filePath, 0644)
	}
	filePath = filepath.Join(filePath, fileName)
	err := sh.Get(hash, filePath)
	d.OnError(err)

}
