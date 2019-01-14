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

// FriendList lists out all your added friends
func FriendList() *cobra.Command {
	var friendlistCmd = &cobra.Command{
		Use:   "friendlist",
		Short: "list out all your added friends",
		Long: `List out all your friends if you have added any, or else return nothing

Example:
	` + tools.Args0() + ` friendlist`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("friendlist called")
		},
	}
	return friendlistCmd
}
