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
	"path/filepath"

	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/spf13/cobra"
)

// UpFile represents the UpFile command
func UpFile() *cobra.Command {
	var upCmd = &cobra.Command{
		Use:   "up",
		Short: "Add and provide your file or directory to the network",
		Long: `Add the file or directory to your local node and provide (publish) the file
to the network so that other nodes are able to retrieve it.

Examples:
	
	` + tools.Args0() + ` up Example.pdf`,
		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			up(cmd, args)
		},
	}

	return upCmd
}

func up(cmd *cobra.Command, args []string) {
	fn := args[0]
	var hash string
	sh := NewIpfsAPI()
	currentDir, err := os.Getwd()
	d.OnError(err)
	TheDir := filepath.Join(currentDir, fn)
	fInfo, err := os.Stat(TheDir)
	d.OnError(err)

	if fInfo.IsDir() {
		hash, err = sh.AddDir(TheDir)
		d.OnError(err)

	} else {
		file, err := os.Open(fn)
		d.OnError(err)
		hash, err = sh.Add(file)
		d.OnError(err)
	}

	fmt.Printf("\n%s is up! with hash %s\n\n", fn, hash)
	fmt.Printf("Peers are able to retrieve the file by:\n\n%s retobject %s <friend's name> -o %s\n",
		tools.Args0(), hash, fn)
}

// AddFile just add local file to local node and pin it
// func AddFile(node *core.IpfsNode, file string) iface.Path {
// 	ctx, cancel := context.WithTimeout(context.Background(), cnf.Timeout)
// 	defer cancel()
// 	f, err := os.Open(file)
// 	d.OnError(err)

// 	hash, err := coreunix.AddWithContext(ctx, node, f)
// 	d.OnError(err)

// 	hashPath, err := iface.ParsePath(hash)
// 	d.OnError(err)

// 	return hashPath
// }
