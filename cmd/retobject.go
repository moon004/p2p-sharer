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
	"fmt"
	"io"
	"os"
	"path/filepath"

	pstore "gx/ipfs/QmPiemjiKBC9VA7vZF82m4x1oygtg2c2YVqag8PX7dN1BD/go-libp2p-peerstore"

	"github.com/ipfs/go-ipfs/core/coreapi"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/moon004/p2p-sharer/friend"
	"github.com/moon004/p2p-sharer/ipfs"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetObject() *cobra.Command {
	var getobject = &cobra.Command{
		Use:   "retobject",
		Short: "Retrieve file from peers",
		Long: `Retrieve file from peers and save to local directory

Examples:
	
	` + tools.Args0() + ` retobject QmSgc9oPMqBppGyM3TWc7NZF11bwH8o3CDekd6pAYGJF8X -p <Peers Info>`,

		Args: cobra.ExactArgs(1),

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
	getobject.Flags().StringP("friendName", "n", "", "Receiver's ID (required)")
	getobject.Flags().StringP("fileName", "f", "", "The output filename (required)")
	return getobject
}

func retobject(cmd *cobra.Command, args []string) {
	friendName, _ := cmd.Flags().GetString("friendName")
	fileName, _ := cmd.Flags().GetString("fileName")
	hash := args[0]
	fmt.Println(friendName, hash)

	node, cancel := NewNodeLoader()
	defer cancel() // cancel the ctx after operation is done
	nodeCtx := node.Context()
	/*
		1. if got provide friendName, find friendName in config file
		2. if provided friendName but cant find == Not Provided == Dont Have
		3. if Dont Have, just dht findprovs, and connect and get <hash>
		4. If findprovs empty, just get <hash>
		5. if Have PeerID, connect and get <hash>
	*/
	var PeerInfo pstore.PeerInfo
	// if -n is not empty
	if friendName != "" {
		// Acquire the friend's Peer's ID
		var f friend.FList
		Flist := f.GetFList()
		PeerInfo = Flist.Friends[friendName]

		if PeerInfo.ID == "" {
			// Induce an error message
			err := errors.Errorf("You have no such friend: %s", friendName)
			d.OnError(err) // stop program here and show error
		}
		// Do step 5.
		api, err := coreapi.NewCoreAPI(node)
		d.OnError(err)

		err = api.Swarm().Connect(nodeCtx, PeerInfo)
		d.OnError(err)

		reader, _, err := ipfs.Cat(nodeCtx, api, hash, 0, -1)
		d.OnError(err)

		// Output to .p2p-sharer/storage
		p2pPath, _ := viper.Get("p2p_config_file").(string)
		dir, _ := filepath.Split(p2pPath)
		filePath := filepath.Join(dir, fileName)
		tmpFile, err := os.Create(filePath)
		d.OnError(err)

		size, err := io.Copy(tmpFile, reader)
		fmt.Printf("Size: %vbyte", size)
	}

	// path, err := iface.ParsePath(hash)
	// d.OnError(err)

	// commands.cat(ctx, api, )
}
