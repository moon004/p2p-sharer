package friend

import (
	"io/ioutil"

	"github.com/moon004/p2p-sharer/cnf"
	"github.com/moon004/p2p-sharer/tools"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// FList is the friend list struct
type FList struct {
	Friends cnf.FriendList `yaml:"friend_list"`
}

// GetFList returns the FList struct
func (f *FList) GetFList() *FList {
	p2pPath := viper.Get("p2p_config_file").(string)
	yamlContent, err := ioutil.ReadFile(p2pPath)
	tools.OnError(err)
	err = yaml.Unmarshal(yamlContent, f)
	tools.OnError(err)
	return f
}
