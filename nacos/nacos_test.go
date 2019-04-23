package nacos

import (
	"fmt"
	"testing"
	"github.com/c-jiazheng/yzf-common/net"
)

func TestNewDiscoveryError(t *testing.T) {
	ip := net.GetFirstIpAddress()
	fmt.Println(ip)
}
