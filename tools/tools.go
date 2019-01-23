package tools

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/moon004/p2p-sharer/ipfs"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Args0 returns string of the binary name
func Args0() string {
	name, err := os.Executable()
	if err != nil {
		return filepath.Base(os.Args[0])
	}

	link, err := filepath.EvalSymlinks(name)
	if err != nil {
		return filepath.Base(name)
	}
	BaseLink := filepath.Base(link)
	pair := strings.Split(BaseLink, ".")
	if len(pair) >= 2 {
		BaseLink = pair[0]
	}

	return BaseLink
}

// OnError handler, display error with stack trace on debug
// while display normal error without debug mode
func OnError(err error) {
	debug := viper.GetBool("debug")
	if err != nil {
		if debug {
			red := color.New(color.FgRed).SprintFunc()
			log.Print(fmt.Sprintf("%s %v", red("error:"), err))
			err = errors.WithStack(err)
			fmt.Printf("%+v\n", err)
		} else {
			red := color.New(color.FgRed).SprintFunc()
			log.Print(fmt.Sprintf("%s %v", red("error:"), err))
		}
		os.Exit(1)
	}
}

// Info for debugging purpose
func Info(v ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	msg := fmt.Sprintln(v...)
	fmt.Print(fmt.Sprintf("%s %s", green("info:"), msg))
}

// NewNodeLoader returns an ipfs node and the context cancel function
func NewNodeLoader() (*core.IpfsNode, context.CancelFunc) {
	dur := GetTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), dur)

	// Invoke LoadPlugins to load plugins into our repo
	//			"" means load New Plugins
	_, err := ipfs.LoadPlugins("")
	OnError(err)

	configPath := viper.Get("ipfs_config_path").(string)
	repo, err := fsrepo.Open(configPath)
	OnError(err)
	cfg := &core.BuildCfg{
		Repo: repo,
	}
	node, err := core.NewNode(ctx, cfg)
	OnError(err)

	return node, cancel

}

// GetTimeout returns the timeout of the of the command
func GetTimeout() time.Duration {
	return time.Duration(viper.GetInt64("timeout")) * time.Second
}
