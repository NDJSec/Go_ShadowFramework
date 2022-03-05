package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	UserInput := bufio.NewScanner(os.Stdin)
	var running bool = true

	for running {
		fmt.Print("SFconsole > ")
		UserInput.Scan()

		switch strings.ToLower(UserInput.Text()) {
		case "scan":
		case "exit":
			break
		}
	}

}
