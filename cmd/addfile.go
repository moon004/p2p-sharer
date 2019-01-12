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

// AddFile represents the addfile command
func AddFile() *cobra.Command {
	var addfileCmd = &cobra.Command{
		Use:   "addfile",
		Short: "Add a file or folder to local node",
		Long: `Add a local file or folder to your local node
		
Examples:
	
	` + tools.Args0() + ` MyFiles.pdf`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("addfile called")
		},
	}

	return addfileCmd
}
