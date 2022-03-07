package portscan

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

type Options struct {
	port string
}

func init() {
	var rhost string
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

func PortScanHandler(UserInput string) {
	switch strings.ToLower(UserInput) {
	case "show options":
		fmt.Println("Module Options (auxiliary/portscan)")
		fmt.Println("Name 	Current Setting	Required	Description")
		fmt.Println("----   --------------- --------    -----------")
		fmt.Fprintln("RHOST  %s				yes			Host to be scanned", rhost)
	case "set":
		//Handle setting options here
	case "help":
		//Print help menu for this module
	}

}

func setPortScanOptions() {

	/*options := &Options{
		ip:
	}*/
}

func portScan() {
	ps := &PortScanner{
		ip:   options.ip,
		lock: semaphore.NewWeighted(Ulimit()),
	}
	ps.Start(1, 65535, 500*time.Millisecond)
}
