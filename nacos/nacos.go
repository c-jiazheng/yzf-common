package nacos

import (
	mynet "github.com/c-jiazheng/yzf-common/net"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/clients/service_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
	"strconv"
	"strings"
)

func RegistryNacosServer(nacosHost,listenAddress,nacosDiscoverClient string) (err error) {

	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	HostPort := strings.Split(nacosHost, ":")
	sListenAddress := strings.Split(listenAddress,":")


	port, _ := strconv.ParseInt(HostPort[1], 10, 64)
	bindPort,_ := strconv.ParseInt(sListenAddress[1],10,64)
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: HostPort[0],
		Port:   uint64(port), //8848,
	}})

	if len(nacosDiscoverClient) == 0 {
		nacosDiscoverClient = mynet.GetFirstIpAddress()
	}

	err = client.StartBeatTask(vo.BeatTaskParam{
		Ip: nacosDiscoverClient,
		Port:    uint64(bindPort),
		//Cluster: "a",
		Dom: "node-exporter",
	})

	return err
}
