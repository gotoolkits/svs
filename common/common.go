package common

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

type Services struct {
	Srvmnt string
	Ss     []Service
}

type Service struct {
	HostName    string
	ServiceName string
	Status      bool
	LastSync    time.Time
	MaxFail     int
	TimeOut     time.Time
	Mnt         int
	Control     bool
}

func (ss *Services) MntIsRuning() (bool, error) {

	cmd := exec.Command("service", ss.Srvmnt, "status")

	d, err := Execute(cmd)

	if err != nil {
		return false, err
	}
	status := strings.Contains(string(d), "running")
	return status, nil

}

func (ss *Services) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
