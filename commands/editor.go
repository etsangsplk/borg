package commands

import (
	"errors"

	"github.com/ok-borg/borg/conf"
)

func init() {
	var summary string = "Editor summary"
	Commands["editor"] = Command{
		F:       Editor,
		Summary: summary,
	}
}

// Login saves a token acquired from the web page into the user config file
func Editor(args []string) error {
	if len(args) < 2 {
		return errors.New("Please supply an editor. The default is vim, so if you are happy with that, do nothing.")
	}
	editor := args[1]
	conf, err := conf.Get()
	if err != nil {
		return err
	}
	conf.Editor = editor
	return conf.Save()
}
