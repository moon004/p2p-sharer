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
	"context"
	"os"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/coreapi/interface"
	"github.com/ipfs/go-ipfs/core/coreunix"
	"github.com/moon004/p2p-sharer/cnf"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// UpFile represents the UpFile command
func UpFile() *cobra.Command {
	var upfileCmd = &cobra.Command{
		Use:   "upfile",
		Short: "Add and provide your file to the network",
		Long: `Add the file to your local node and provide (publish) the file
to the network so that other nodes are able to retrieve it.

Examples:
	
	` + tools.Args0() + ` -f Example.pdf`,
		Run: func(cmd *cobra.Command, args []string) {
			allflags := cmd.Flags()
			if allflags.Changed("filename") == false {
				tools.OnError(errors.New("Must provide value for all the required flag"))
				return
			}
			upfile(cmd, args)
		},
	}

	upfileCmd.Flags().SortFlags = false
	upfileCmd.Flags().StringP("filename", "f", "", "Name of the file to upload (required)")
	return upfileCmd
}

func upfile(cmd *cobra.Command, args []string) {
	fn, _ := cmd.Flags().GetString("filename")

	node, cancel := tools.NewNodeLoader()
	defer cancel()
	nodeCtx := node.Context()

	hashPath := AddFile(node, fn)

	api, err := coreapi.NewCoreAPI(node)
	tools.OnError(err)

	err = api.Dht().Provide(nodeCtx, hashPath)
	tools.OnError(err)

}

// AddFile just add local file to local node and pin it
func AddFile(node *core.IpfsNode, file string) iface.Path {
	ctx, cancel := context.WithTimeout(context.Background(), cnf.Timeout)
	defer cancel()
	f, err := os.Open(file)
	tools.OnError(err)

	hash, err := coreunix.AddWithContext(ctx, node, f)
	tools.OnError(err)

	hashPath, err := iface.ParsePath(hash)
	tools.OnError(err)

	return hashPath
}
