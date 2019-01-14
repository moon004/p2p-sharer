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
	"os"

	"github.com/moon004/p2p-sharer/tools"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SendFile represents the sendfile command
func SendFile() *cobra.Command {
	var sendfileCmd = &cobra.Command{
		Use:   "sendfile",
		Short: "Send your file to your peers",
		Long: `Send local file to your peers via peers ID

Examples:
	
	` + tools.Args0() + ` -f Example.pdf -p 'Peers ID'`,
		Run: func(cmd *cobra.Command, args []string) {
			allflags := cmd.Flags()
			if allflags.Changed("peerID") == false ||
				allflags.Changed("filename") == false {

				tools.OnError(errors.New("Must provide value for all the required flag"))
				return
			}
			sendfile(cmd, args)
		},
	}

	sendfileCmd.Flags().SortFlags = false
	sendfileCmd.Flags().StringP("peerID", "p", "", "Receiver's ID (required)")
	sendfileCmd.Flags().StringP("filename", "f", "", "Name of the file to send (required)")
	return sendfileCmd
}

func sendfile(cmd *cobra.Command, args []string) {
	ID, _ := cmd.Flags().GetString("peerID")
	fn, _ := cmd.Flags().GetString("filename")

	// Add the localfile to ipfs
	f, err := os.Open(fn)
	if err != nil {
		tools.OnError(err)
	}

	fmt.Println("f", f, ID)
}
