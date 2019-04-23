package nacos

import (
	"fmt"
	"net"

	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/clients/service_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
	"os"
	"strconv"
	"strings"
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
				ip = ipnet.IP.String()
				break
			}

		}
	}
	return
}

func RegistryNacosServer(nacosHost, listenAddress, nacosDiscoverClient string, nodeType string) (err error) {

	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	HostPort := strings.Split(nacosHost, ":")
	sListenAddress := strings.Split(listenAddress, ":")

	port, _ := strconv.ParseInt(HostPort[1], 10, 64)
	bindPort, _ := strconv.ParseInt(sListenAddress[1], 10, 64)
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: HostPort[0],
		Port:   uint64(port), //8848,
	}})

	if len(nacosDiscoverClient) == 0 {
		nacosDiscoverClient = GetFirstIpAddress()
	}

	/*success, _ := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          nacosDiscoverClient,
		Port:        uint64(port),
		ServiceName: "node-exporter",
		Weight:      1000,
		//ClusterName: "a",
		Metadata: map[string]string{"node-type":nodeType},
	})
	fmt.Println(success)*/

	err = client.StartBeatTask(vo.BeatTaskParam{
		Ip:   nacosDiscoverClient,
		Port: uint64(bindPort),
		//Cluster: "a",
		Dom:      "node-exporter",
		Metadata: map[string]string{"node-type": nodeType},
	})

	return err
}
