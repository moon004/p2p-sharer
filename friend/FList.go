package friend

import (
	"io/ioutil"

	"github.com/moon004/p2p-sharer/cnf"
	d "github.com/moon004/p2p-sharer/debugs"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// FList is the friend list struct
type FList struct {
	Friends cnf.FriendList `yaml:"friend_list"`
}

// GetFList returns the FList struct
func (f *FList) GetFList() *FList {
	p2pPath := viper.Get("friend_list").(string)
	yamlContent, err := ioutil.ReadFile(p2pPath)
	d.OnError(err)
	err = yaml.Unmarshal(yamlContent, f)
	d.OnError(err)
	return f
}
