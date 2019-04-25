## prometheus指标采集

[![Build Status](https://cloud.drone.io/api/badges/oliver006/redis_exporter/status.svg)](https://cloud.drone.io/oliver006/redis_exporter)
 [![Coverage Status](https://coveralls.io/repos/github/oliver006/redis_exporter/badge.svg?branch=master)](https://coveralls.io/github/oliver006/redis_exporter?branch=master) 
 [![codecov](https://codecov.io/gh/oliver006/redis_exporter/branch/master/graph/badge.svg)](https://codecov.io/gh/oliver006/redis_exporter)
 
 ### 参数
 - nacos.server-address nacos 注册中心地址\
 默认 172.24.28.2:8848,可以设置evn NACOS_SERVER_ADDRESS
 - nacos.discover-client-address \
 注册到nacos的服务的客户端地址，evn NACOS_DISCOVER_CLIENT_ADDRESS,如果不设置则取第一个网卡地址。
 - nacos.node-type \
 注册到nacos的节点类型 默认为：exporter
 - nacos.service-name \
 注册到nacos的服务明