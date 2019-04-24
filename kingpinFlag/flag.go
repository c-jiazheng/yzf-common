package kingpinFlag

import "gopkg.in/alecthomas/kingpin.v2"

//添加参数定义
func InitFlag() (nacosHost, nacosDiscoverClient, nodeType, serviceName *string) {

	//nacos注册中心地址
	nacosHost = kingpin.Flag(
		"nacos.host",
		"Nacos Host default http://172.24.28.2:8848",
	).Default("172.24.28.2:8848").Envar("NACOS_HOST").String()
	//注册到nacos的服务地址
	nacosDiscoverClient = kingpin.Flag(
		"nacos.discoverClient.ip",
		"Registry Nacos client ip address.",
	).Default("").Envar("NACOS_DISCOVER_CLIENT_IP").String()
	//节点类型
	nodeType = kingpin.Flag(
		"nacos.nodeType",
		"Registry Nacos meteData node_type.",
	).Default("exporter").String()
	//服务名称
	serviceName = kingpin.Flag(
		"nacos.serviceName",
		"Registry Nacos serviceName.",
	).Default("node-exporter").String()

	return
}
