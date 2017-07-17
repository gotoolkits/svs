package common

import (
	"regexp"
	"strings"
)

/*
##Samples:##
calico-felix.service                  disabled
calico-node.service                   disabled
*/
func SysEnExtract(d []byte) []string {

	var srvNameList []string

	re := regexp.MustCompile(".*\n")
	str := re.FindAllString(string(d), -1)

	// 拿掉第一行和最后两行
	str = str[1 : len(str)-2]

	// 获取enabled的服务列表
	for _, v := range str {

		tmp := strings.Fields(v)
		//s[tmp[0]] = tmp[1]

		if tmp[1] == "enabled" {

			srvNameList = append(srvNameList, tmp[0])

		}
	}
	//	fmt.Println(len(srvNameList))
	return srvNameList
}

/*
##Samples:##
  Active: active (running)
  Active: inactive (dead)
*/
func SysRnExtract(d []byte) string {

	re := regexp.MustCompile("Active:.*\n")
	str := re.FindString(string(d))

	return str
}

func SupRnExtract(d []byte) string {
	return ""

}
