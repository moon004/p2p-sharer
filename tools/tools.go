package tools

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

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

// GetTimeout returns the timeout of the of the command
func GetTimeout() time.Duration {
	return time.Duration(viper.GetInt64("timeout")) * time.Second
}

func Regex(option string) (*regexp.Regexp, error) {

	switch option {
	case "ipfs hash":
		return regexp.MustCompile(`Qm[a-zA-Z0-9]{44}`), nil
	case "string input":
		return regexp.MustCompile(`^\w[a-zA-Z0-9-_\s]+\w$`), nil
	case "ipfs address":
		return regexp.MustCompile(
			`\/ip4\/\d+\.\d+\.\d+\.\d+\/tcp\/\d+\/ipfs\/Qm[a-zA-Z0-9]{44}`,
		), nil
	default:
		return nil, errors.New("invalid option")
	}
}
