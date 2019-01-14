package cnf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// ConfigStruct the structure of Config File (yaml)
type ConfigStruct struct {
	UserLocalID string `yaml:"local_id"`
	Version     string `yaml:"version"`
	Verbose     bool   `yaml:"verbose"`
	Debug       bool   `yaml:"debug"`
}

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
	return filepath.Join(c.ConfigDir(), ConfigFileName)
}

// ConfigDir is the directory of the config file
func (c *ConfigStruct) ConfigDir() string {
	p, err := c.Path()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	dirname := tools.Args0()

	return filepath.Join(p, "."+dirname)
}

// Path return the home directory of the executor
func (c *ConfigStruct) Path() (string, error) {
	p, err := homedir.Dir()
	if err != nil {
		return "", errors.Wrap(err, "failed to detect homeDir")
	}
	return p, nil
}

func (c *ConfigStruct) WriteToConfig() error {
	// Write to the config yaml file
	content, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "WriteToConfig marshal failed")
	}

	if !exist(c.ConfigDir()) {
		_ = os.Mkdir(c.ConfigDir(), os.ModePerm)
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
	defaultcnf := ConfigStruct{
		Version:     BaseVersion,
		Verbose:     false,
		Debug:       false,
		UserLocalID: c.GetLocalIPFSID(),
	}
	err := defaultcnf.WriteToConfig()
	err = errors.Wrap(err, "error writing default value to config")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	return viper.ReadInConfig()
}

type ID struct {
	PeerID string `json:"PeerID"`
}

type IDRetriever struct {
	Identity *ID `json:"Identity"`
}

// GetLocalIPFSID is to get the local node ID
func (c *ConfigStruct) GetLocalIPFSID() string {
	p, _ := c.Path()
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

	return Retriever.Identity.PeerID
}
