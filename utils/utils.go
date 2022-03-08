package utils

var runningPortScan, runningWinRevShellClient, runningWinRevShellServer bool

func SetRunVar(runtype string, status bool) {
	switch runtype {
	case "portscan":
		runningPortScan = status
	case "winrevshell_client":
		runningWinRevShellClient = status
	case "winrevshell_handler":
		runningWinRevShellServer = status
	}

}

func GetRunVar(runtype string) bool {
	switch runtype {
	case "portscan":
		return runningPortScan
	case "winrevshell_client":
		return runningWinRevShellClient
	case "winrevshell_handler":
		return runningWinRevShellServer
	}

	return false
}
