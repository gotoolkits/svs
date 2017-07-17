package modules

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/svs/common"
)

type Systemd struct {
	IsExist   bool
	IsRunning bool
}

func (sys *Systemd) GetList() ([]string, error) {
	var cmd *exec.Cmd

	cmd = exec.Command("systemctl", "list-unit-files", "--type=service")

	d, err := common.Execute(cmd)
	if err != nil {
		return nil, err
	}

	return common.SysEnExtract(d), err

}

func (sys *Systemd) IsRuning(s string) (bool, error) {

	var cmd *exec.Cmd

	cmd = exec.Command("systemctl", "status", s)

	d, err := common.Execute(cmd)

	if err != nil {
		return false, err
	}

	str := common.SysRnExtract(d)

	// fmt.Println(str)

	running := strings.Contains(str, "running")
	dead := strings.Contains(str, "dead")

	if running {
		return true, nil
	}

	if dead {
		return false, nil
	}

	return false, nil

}

func (sys *Systemd) Run() {

	status, err := SrvList.MntIsRuning()
	if err != nil {
		fmt.Println("management err")
	}

	fmt.Println(SrvList.Srvmnt, "running status:", status)

	// 获取systemd服务enabled的列表
	list, err := Sys.GetList()
	if err != nil {
		fmt.Println("getlist error")
	}

	// fmt.Println(list)

	for _, v := range list {
		Srv.HostName = "127.0.0.1"
		Srv.ServiceName = v
		Srv.Status, _ = Sys.IsRuning(v)
		Srv.LastSync = time.Now()
		Srv.Mnt = 1

		SrvList.Ss = append(SrvList.Ss, Srv)

	}

	for i := 0; i < len(SrvList.Ss); i++ {

		fmt.Printf("ServiceName:%s \t\t is running:%v \n", SrvList.Ss[i].ServiceName, SrvList.Ss[i].Status)

	}
}
