package utils

var runningPortScan bool

func SetRunVar(runtype string, status bool) {
	switch runtype {
	case "portscan":
		runningPortScan = status
	}
}

func GetRunVar(runtype string) bool {
	switch runtype {
	case "portscan":
		return runningPortScan
	}

	return false
}
