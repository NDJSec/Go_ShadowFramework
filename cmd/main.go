package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDJSec/Go_ShadowFramework/modules/auxiliary/portscan"
	"github.com/NDJSec/Go_ShadowFramework/utils"
)

func main() {
	UserInput := bufio.NewScanner(os.Stdin)
	var running bool = true

	for running {
		fmt.Print("SFconsole > ")
		UserInput.Scan()

		switch strings.ToLower(UserInput.Text()) {
		case "scan":
			fmt.Print("SFconsole auxiliary(portscan) > ")
			UserInput.Scan()
			utils.SetRunVar("portscan", true)
			for utils.GetRunVar("portscan") {
				portscan.PortScanHandler(UserInput.Text())
				if UserInput.Text() != "back" {
					fmt.Print("SFconsole auxiliary(portscan) > ")
					UserInput.Scan()
				}

			}
		case "exit":
			running = false
		}
	}

}
