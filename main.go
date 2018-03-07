package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sethdmoore/lemonfeeder/constants"
	"github.com/sethdmoore/lemonfeeder/types"
	"github.com/sethdmoore/lemonfeeder/utils"
	"os"
	"strings"
)

func main() {
	// cmd, err := utils.Popen("ls", "-al", "/")
	cmd, err := utils.Popen("printf", "")
	var output string
	var icons []string
	if err != nil {
		fmt.Printf("OH NO: %s\n", err)
	}

	conf, err := types.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize config: %s\n", err)
	}
	//fmt.Printf(fmt.Sprintf("%v, %v", conf, conf))
	spew.Fdump(os.Stderr, conf)

	for i := 0; i < conf.DesktopCount; i++ {
		//strings.Join()
		icons = append(icons, constants.Icons[fmt.Sprintf("desktop%d", i+1)])
	}

	for {
		output = ""
		//output += constants.Icons[fmt.Sprintf("desktop%d", i+1)] + "   "

		output = strings.Join(icons, "   ")

		fmt.Printf(cmd)
		fmt.Printf(output)
		fmt.Printf("\n")
		utils.Sleep(1)
	}
}
