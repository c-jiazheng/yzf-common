package kingpinFlag

import "gopkg.in/alecthomas/kingpin.v2"

//添加参数定义
func InitFlag(serviceName string) (nacosHost, nacosDiscoverClient, nodeType *string) {

	//nacos注册中心地址
	nacosHost = kingpin.Flag(
		"nacos.server-address",
		"Nacos server address default http://172.24.28.2:8848",
	).Default("172.24.28.2:8848").Envar("NACOS_SERVER_ADDRESS").String()
	//注册到nacos的服务地址
	nacosDiscoverClient = kingpin.Flag(
		"nacos.discover-client-address",
		"Registry Nacos client address.",
	).Default("").Envar("NACOS_DISCOVER_CLIENT_ADDRESS").String()
	//节点类型
	nodeType = kingpin.Flag(
		"nacos.node-type",
		"Registry Nacos meteData node_type.",
	).Default("exporter").String()
	//服务名称
	kingpin.Flag(
		"nacos.service-name",
		"Registry Nacos serviceName.",
	).Default(serviceName).String()

	return
}
