package cnf

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	pstore "gx/ipfs/QmPiemjiKBC9VA7vZF82m4x1oygtg2c2YVqag8PX7dN1BD/go-libp2p-peerstore"
	"gx/ipfs/QmWGm4AbZEbnmdgVTza52MSNpEmBdFVqzmAysRbjrRyGbH/go-ipfs-cmds"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	homedir "github.com/mitchellh/go-homedir"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/moon004/p2p-sharer/ipfs"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// ConfigStruct the structure of Config File (yaml)
type ConfigStruct struct {
	UserLocalID *pstore.PeerInfo `yaml:"local_id"`
	IpfsConFile string           `yaml:"ipfs_config_path"`
	P2pConFile  string           `yaml:"p2p_config_file"`
	Version     string           `yaml:"version"`
	Friends     []FriendList     `yaml:"friend_list"`
	Verbose     bool             `yaml:"verbose"`
	Debug       bool             `yaml:"debug"`
}

type FriendList map[string]pstore.PeerInfo

// Reload Read the config file and Unmarshal
// into the ConfigStruct
func (c *ConfigStruct) Reload() error {
	content, err := ioutil.ReadFile(c.ConfigFile())
	if err != nil {
		return errors.Wrap(err, "Readfile failed")
	}

	if err = yaml.Unmarshal(content, c); err != nil {
		return errors.Wrap(err, "Unmarshal Failed")
	}

	return nil
}

// ConfigFile is the EXACT directory of the file
func (c *ConfigStruct) ConfigFile() string {
	//																	config.yaml
	return filepath.Join(ConfigDir(), ConfigFileName)
}

// ConfigDir is the directory of the config file of p2p-sharer
func ConfigDir() string {
	p, err := Path()
	d.OnError(err)

	dirname := tools.Args0()

	return filepath.Join(p, "."+dirname)
}

// IpfsConfDir returns the Ipfs root directory
// Initialize an ipfs config if it doesn't have one
func IpfsConfDir() string {
	var req *cmds.Request

	path, err := ipfs.GetRepoPath(req)
	d.OnError(err)

	if !fsrepo.IsInitialized(path) {
		// Init the ipfs config and repo
		err := ipfs.InitWithDefaults(os.Stdout, path, "")
		d.OnError(err)
	}

	return path
}

// Path return the home directory of the executor
func Path() (string, error) {
	p, err := homedir.Dir()
	if err != nil {
		return "", errors.Wrap(err, "failed to detect homeDir")
	}
	return p, nil
}

// WriteToConfig update the p2p-sharer config file
func (c *ConfigStruct) WriteToConfig() error {
	// Write to the config yaml file
	content, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "WriteToConfig marshal failed")
	}

	if !exist(ConfigDir()) {
		_ = os.Mkdir(ConfigDir(), os.ModePerm)
	}

	return ioutil.WriteFile(c.ConfigFile(), content, 0644)
}

func exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadConFile update the struct with current config file value
func ReadConFile() *ConfigStruct {
	cnf := &ConfigStruct{}
	err := cnf.Reload()
	if err != nil {
		return nil
	}

	return cnf
}

// DefaultConfigValue update the config file with default value
func (c *ConfigStruct) DefaultConfigValue() error {
	ipfsFilePath := IpfsConfDir()
	defaultcnf := ConfigStruct{
		Version:     BaseVersion,
		Verbose:     false,
		Debug:       false,
		Friends:     make([]FriendList, 0),
		IpfsConFile: ipfsFilePath,
		P2pConFile:  c.ConfigFile(),
		UserLocalID: GetLocalIPFSID(),
	}
	err := defaultcnf.WriteToConfig()
	err = errors.Wrap(err, "error writing default value to config")
	d.OnError(err)

	return viper.ReadInConfig()
}

type ID struct {
	PeerID string `json:"PeerID"`
}

type IDRetriever struct {
	Identity *ID `json:"Identity"`
}

// GetLocalIPFSID is to get the local node ID
func GetLocalIPFSID() *pstore.PeerInfo {
	p, _ := Path()
	dir := filepath.Join(p, ".ipfs", "config")
	jsonData, err := ioutil.ReadFile(dir)
	if err != nil {
		err = errors.Wrap(err, "error Reading ipfs json config file")
		log.Fatalf("%+v", err)
	}
	// Retrieve ipfs local id
	Retriever := &IDRetriever{}
	err = json.Unmarshal(jsonData, &Retriever)
	if err != nil {
		err = errors.Wrap(err, "error marshalling ID into json")
		log.Fatalf("%+v", err)
	}
	// Getting the swarm addrs local id
	node, cancel := NewNodeInit()
	defer cancel()
	nodeCtx := node.Context()
	api, err := coreapi.NewCoreAPI(node)
	d.OnError(err)

	maddrs, err := api.Swarm().LocalAddrs(nodeCtx)
	d.OnError(err)

	PeerInfo, err := pstore.InfoFromP2pAddr(maddrs[0])
	d.OnError(err)

	return PeerInfo
}

// NewNodeInit initialize NewNode when creating config file
func NewNodeInit() (*core.IpfsNode, context.CancelFunc) {
	dur := tools.GetTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), dur)

	// Invoke LoadPlugins to load plugins into our repo
	//			"" means load New Plugins
	_, err := ipfs.LoadPlugins("")
	d.OnError(err)

	configPath := IpfsConfDir()
	repo, err := fsrepo.Open(configPath)
	d.OnError(err)
	cfg := &core.BuildCfg{
		Repo: repo,
	}
	node, err := core.NewNode(ctx, cfg)
	d.OnError(err)

	return node, cancel
}
