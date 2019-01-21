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

	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetObject() *cobra.Command {
	var getobject = &cobra.Command{
		Use:   "retobject",
		Short: "Retrieve file from peers",
		Long: `Retrieve file from peers and save to local directory

Examples:
	
	` + tools.Args0() + ` QmSgc9oPMqBppGyM3TWc7NZF11bwH8o3CDekd6pAYGJF8X -p 'Peers ID'`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			allflags := cmd.Flags()
			if allflags.Changed("peerID") == false {
				tools.OnError(errors.New("Must provide value for all the required flag"))
				return
			}
			retobject(cmd, args)
		},
	}

	getobject.Flags().SortFlags = false
	getobject.Flags().StringP("peerID", "p", "", "Receiver's ID (required)")
	return getobject
}

func retobject(cmd *cobra.Command, args []string) {
	ID, _ := cmd.Flags().GetString("peerID")
	hash := args[0]
	fmt.Println(ID, hash)
	// path, err := iface.ParsePath(hash)
	// tools.OnError(err)

	// commands.cat(ctx, api, )
}
