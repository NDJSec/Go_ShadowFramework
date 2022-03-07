package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDJSec/Go_ShadowFramework/modules/auxiliary/portscan"
<<<<<<< HEAD
	"github.com/NDJSec/Go_ShadowFramework/utils"
=======
>>>>>>> 9d1fbace483c3c0ffe5903570b53fb1fc2f2303e
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
<<<<<<< HEAD
			utils.SetRunVar("portscan", true)
			for utils.GetRunVar("portscan") {
				portscan.PortScanHandler(UserInput.Text())
				if UserInput.Text() != "back" {
					fmt.Print("SFconsole auxiliary(portscan) > ")
					UserInput.Scan()
				}

			}
=======
			portscan.PortScanHandler(UserInput.Text())
>>>>>>> 9d1fbace483c3c0ffe5903570b53fb1fc2f2303e
		case "exit":
			running = false
		}
	}

}
