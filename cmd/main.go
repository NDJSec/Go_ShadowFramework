package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDJSec/Go_ShadowFramework/modules/auxiliary/portscan"
	winrevshell "github.com/NDJSec/Go_ShadowFramework/modules/exploit/win_rev_shell"
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
		case "test rev shell":
			fmt.Print("SFconsole exploit(win_rev_shell) > ")
			UserInput.Scan()
			utils.SetRunVar("winrevshell_client", true)
			for utils.GetRunVar("winrevshell_client") {
				portscan.PortScanHandler(UserInput.Text())
				if UserInput.Text() != "back" {
					fmt.Print("SFconsole exploit(win_rev_shell) > ")
					UserInput.Scan()
				}

			}
		case "winshell":
			fmt.Print("SFconsole exploit(win_rev_shell_handler) > ")
			UserInput.Scan()
			utils.SetRunVar("winrevshell_handler", true)
			for utils.GetRunVar("winrevshell_handler") {
				winrevshell.WinRevShellServerHandler(UserInput.Text())
				if UserInput.Text() != "back" {
					fmt.Print("SFconsole exploit(win_rev_shell_handler) > ")
					UserInput.Scan()
				}

			}

		case "exit":
			running = false
		}
	}

}
