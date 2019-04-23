package nacos

import (
	"fmt"
	"testing"
	"yzf-common/net"
)

func TestNewDiscoveryError(t *testing.T) {
	ip := net.GetFirstIpAddress()
	fmt.Println(ip)
}
