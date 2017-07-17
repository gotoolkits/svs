package common

import (
	"io/ioutil"
	"os/exec"
)

func Execute(c *exec.Cmd) ([]byte, error) {
	var err error
	var d []byte

	stdout, _ := c.StdoutPipe()
	err = c.Start()
	if err != nil {
		return nil, err
	}

	d, err = ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	err = c.Wait()
	if err != nil {
		return nil, err
	}

	return d, nil

}
