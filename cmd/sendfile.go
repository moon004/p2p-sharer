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

// SendFile represents the sendfile command
func SendFile() *cobra.Command {
	var sendfileCmd = &cobra.Command{
		Use:   "sendfile",
		Short: "Send your file to your peers",
		Long: `Send local file to your peers via peers ID

	Examples:
	
		` + tools.Args0() + ` 'Peers ID'`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sendfile called")
		},
	}

	return sendfileCmd
}
