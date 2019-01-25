package debugs

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

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
