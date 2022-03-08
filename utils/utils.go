package utils

var runningPortScan, runningWinRevShell bool

func SetRunVar(runtype string, status bool) {
	switch runtype {
	case "portscan":
		runningPortScan = status
	case "winrevshell":
		runningWinRevShell = status
	}

}

func GetRunVar(runtype string) bool {
	switch runtype {
	case "portscan":
		return runningPortScan
	case "winrevshell":
		return runningWinRevShell
	}

	return false
}
