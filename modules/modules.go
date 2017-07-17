package modules

import (
	"fmt"

	"github.com/svs/common"
)

const (
	T_SUPERVISOR = "supervisord"
	T_SYSTEMD    = "systemd"
)

type Collector interface {
	GetList() ([]string, error)
	IsRuning(s string) (bool, error)
}

//声明全局变量
var Sys Systemd
var Sup Supervisord
var SrvList common.Services
var Srv common.Service

//模块初始化
func Init(t string) {
	SrvList = common.Services{}
	Srv = common.Service{}

	SrvList.Srvmnt = t

	if t == T_SYSTEMD {
		Sys = Systemd{}
		Run(Sys)
	}

	if t == T_SUPERVISOR {
		Sup = Supervisord{}
		//Sup.Run()
	}

}

func Run(c interface{}) {

	switch v := c.(type) {
	case Systemd:
		Sys.Run()
	case Supervisord:
		Sup.Run()
	default:
		fmt.Println("Unknow type !", v)
	}

}
