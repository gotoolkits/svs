package main

import (
	"flag"
	"fmt"

	"github.com/svs/modules"
)

func main() {

	mType := flag.Int("t", 1, "type1:Systemd | type2:Supervisor")
	hPort := flag.String("p", "9526", "http port")
	maxFail := flag.Int("m", 3, "Max of check service failed times")

	flag.Parse()
	fmt.Println(*mType, *hPort, *maxFail)

	var t string

	if *mType == 1 {
		t = "systemd"
	} else {
		t = "supervisord"
	}

	modules.Init(t)

}
