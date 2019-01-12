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
	"github.com/spf13/cobra"
)

// ConnPeers represents the connpeers comman
func ConnPeers() *cobra.Command {
	var connpeersCmd = &cobra.Command{
		Use:   "connpeers",
		Short: "Connect to an ipfs peer",
		Long: `Connect to a peer so that future request of an ipfs object would be faster.

	Peer's ID can be found using the command 'ipfs id'.

	You can use it like so:

		` + tools.Args0() + ` 'Peer's IPFS ID'

	Examples:

		` + tools.Args0() + ` Qma1dYuhcKgaUP5nooYaHUTQR3phBm6igbbDt9V6Viqo1z`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("connpeers called")
		},
	}

	return connpeersCmd
}
