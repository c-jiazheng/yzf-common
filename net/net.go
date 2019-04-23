package net

import (
	"fmt"
	"net"
	"os"
)

func GetFirstIpAddress() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {

				//log.Debugln("ip:", ipnet.IP.String())
				//append(ips, ipnet.IP.String())
				ip =  ipnet.IP.String()
				break
			}

		}
	}
	return
}
