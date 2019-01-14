package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
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
	}
}
