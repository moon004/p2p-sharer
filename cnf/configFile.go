package cnf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	cmds "gx/ipfs/QmR77mMvvh8mJBBWQmBfQBu8oD38NUN4KE9SL2gDgAQNc6/go-ipfs-cmds"

	api "github.com/ipfs/go-ipfs-api"
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
	UserLocalID *api.IdOutput `yaml:"local_id"`
	IpfsConFile string        `yaml:"ipfs_config_path"`
	P2pConFile  string        `yaml:"p2p_config_file"`
	Version     string        `yaml:"version"`
	Friends     Friend        `yaml:"friend_list"`
	Verbose     bool          `yaml:"verbose"`
	Timeout     uint          `yaml:"timeout"`
	Debug       bool          `yaml:"debug"`
}

// Friend is just a map of string to string
type Friend map[string]string

// Reload Read the config file and Unmarshal
// into the ConfigStruct and return the
func (c *ConfigStruct) Reload() error {
	content, err := ioutil.ReadFile(c.ConfigFile())
	if err != nil {
		return errors.Wrap(err, "Readfile failed")
	}

	if err = yaml.Unmarshal(content, c); err != nil {
		return errors.Wrap(err, "Unmarshal failed")
	}

	return nil
}

//AddFriend adds the new friend new as map[name]ID and save in the config file
func (c *ConfigStruct) AddFriend(name, ID string) error {
	// get the old friend list
	err := c.Reload()
	if err != nil {
		return errors.Wrap(err, "error reloading")
	}
	// add the  new set of NewFriendList
	c.Friends[name] = ID
	err = c.WriteToConfig()
	if err != nil {
		return errors.Wrap(err, "error writing to config")
	}

	return nil
}

// GetFList returns the FList struct
func (c *ConfigStruct) GetFList() (Friend, error) {
	err := c.Reload()
	if err != nil {
		return nil, err
	}

	return c.Friends, nil
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
		Timeout:     120,
		Friends:     make(Friend, 0),
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
func GetLocalIPFSID() *api.IdOutput {
	// Getting the swarm addrs local id
	sh := api.NewLocalShell()
	IDOutput, err := sh.ID()
	d.OnError(err)
	return IDOutput
}
