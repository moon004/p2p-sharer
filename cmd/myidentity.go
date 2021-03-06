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
	"github.com/spf13/viper"
)

// MyIdentity represents the UpFile command
func MyIdentity() *cobra.Command {
	var myidentityCmd = &cobra.Command{
		Use:   "myid",
		Short: "Your local node identity",
		Long: `Share your identity to your other peers so that
they could add you as friend.
Examples:
	` + tools.Args0() + ` myid`,
		Run: func(cmd *cobra.Command, args []string) {
			myidentity()
		},
	}

	myidentityCmd.Flags().SortFlags = false
	myidentityCmd.Flags().StringP("filename", "f", "", "Name of the file to upload (required)")
	return myidentityCmd
}

func myidentity() {
	MyInfo := viper.Get("local_id").(map[string]interface{})
	fmt.Printf("Your IPFS ID: %s\n\n", MyInfo["id"])
	fmt.Println("Your friends/peers can add you as friend via any of these addresses:")
	for _, addrs := range MyInfo["addresses"].([]interface{}) {
		fmt.Printf("%s\n", addrs)
	}
}
