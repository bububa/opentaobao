package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// SlbAliyuncsComDescribeBackendServers20130221 后端服务器健康检查
// slb.aliyuncs.com.DescribeBackendServers.2013-02-21
//
// 后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
func SlbAliyuncsComDescribeBackendServers20130221(clt *core.SDKClient, req *aliyun.SlbAliyuncsComDescribeBackendServers20130221APIRequest, session string) (*aliyun.SlbAliyuncsComDescribeBackendServers20130221APIResponse, error) {
	var resp aliyun.SlbAliyuncsComDescribeBackendServers20130221APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
