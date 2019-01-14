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

// AddFriend adds your ipfs peers as friend
func AddFriend() *cobra.Command {
	var addfriendCmd = &cobra.Command{
		Use:   "addfriend",
		Short: "add your IPFS peers friend",
		Long: `Add your IPFS peers as friend by linking their peer's ID to your own respective naming.

Example:
	` + tools.Args0() + ` Qma1dYuhcKgaUP5nooYaHUTQR3phBm6igbbDt9V6Viqo1z 'Friend's name'`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("addfriend called")
		},
	}
	return addfriendCmd
}
