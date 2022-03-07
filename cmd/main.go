package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDJSec/Go_ShadowFramework/modules/auxiliary/portscan"
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
			portscan.PortScanHandler(UserInput.Text())
		case "exit":
			break
		}
	}

}
