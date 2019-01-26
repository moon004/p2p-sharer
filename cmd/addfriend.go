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
	` + tools.Args0() + ` Qma1dYuhcKgaUP5nooYaHUTQR3phBm6igbbDt9V6Viqo1z 'Friend's name'`,

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

	var friend cnf.Friend
	var cfgStruct cnf.ConfigStruct
	// regex check for name and Addresses
	strInput, _ := tools.Regex("string input")
	ipfsHash, _ := tools.Regex("ipfs address")

	if strInput.MatchString(friendName) && ipfsHash.MatchString(friendID) {
		friend[friendName] = friendID
		err := cfgStruct.AddFriend(friend)
		d.OnError(err)
	} else {
		d.OnError(errors.New("Invalid friend name format or invalid ipfs hash"))
	}
}
