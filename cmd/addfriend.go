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

	"github.com/moon004/p2p-sharer/cnf"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// AddFriend adds your ipfs peers as friend
func AddFriend() *cobra.Command {
	var addfriendCmd = &cobra.Command{
		Use:   "addfriend",
		Short: "add your IPFS peers friend",
		Long: `Add your IPFS peers as friend by linking their peer's ID to your own respective naming.

Example:
	` + tools.Args0() + ` addfriend <peer's identity> <friend's name>`,

		Args: cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			addfriend(cmd, args)
		},
	}
	return addfriendCmd
}

func addfriend(cmd *cobra.Command, args []string) {
	friendID := args[0]
	friendName := args[1]

	fmt.Println(friendID, friendName)
	// regex check for name and Addresses
	strInput, _ := tools.Regex("string input")
	ipfsHash, _ := tools.Regex("ipfs address")

	switch {
	case !strInput.MatchString(friendName):
		d.OnError(errors.New("Invalid friend name format, try other name"))
	case !ipfsHash.MatchString(friendID):
		d.OnError(errors.New("Invalid friend ID"))
	}

	var cfgStruct cnf.ConfigStruct
	err := cfgStruct.AddFriend(friendName, friendID)
	d.OnError(err)
}
