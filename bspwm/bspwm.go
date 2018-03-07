package bspwm

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/sethdmoore/lemonfeeder/constants"
	"github.com/sethdmoore/lemonfeeder/utils"
	"os"
	//"json"
	"strings"
)

func GetDesktops() (int, error) {
	//fmt.Println("vim-go")
	return 3, nil
}

func GetMonitors() ([]string, error) {
	result, err := utils.Popen(constants.BspwmMonitors...)
	if err != nil {
		return nil, err
	}

	if result == "" {
		return nil, errors.New("Empty string from bspc, no monitors?")
	}

	result = result[:len(result)-1]

	spew.Fdump(os.Stderr, result)

	return strings.Split(result, "\n"), nil
}
