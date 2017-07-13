package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sethdmoore/lemonfeeder/config"
	"github.com/sethdmoore/lemonfeeder/constants"
	"github.com/sethdmoore/lemonfeeder/utils"
)

func main() {
	// cmd, err := utils.Popen("ls", "-al", "/")
	cmd, err := utils.Popen("printf", "")
	var output string
	if err != nil {
		fmt.Printf("OH NO: %s\n", err)
	}

	conf := config.New()
	spew.Dump(conf)

	for {
		output = ""
		for i := 1; i < 7; i++ {
			output += constants.Icons[fmt.Sprintf("desktop%d", i)]
		}
		fmt.Printf(cmd.String())
		fmt.Printf(output)
		fmt.Printf("\n")
	}
}
