package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
后端服务器健康检查 APIResponse
slb.aliyuncs.com.DescribeBackendServers.2013-02-21

后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
*/
type SlbAliyuncsComDescribeBackendServers2013-02-21APIResponse struct {
    model.CommonResponse
    Response *SlbAliyuncsComDescribeBackendServers2013-02-21Response `json:"slb_aliyuncs_com_DescribeBackendServers_2013-02-21_response,omitempty"`
}

type SlbAliyuncsComDescribeBackendServers2013-02-21Response struct {

    // request id
    Requestid   string `json:"requestid,omitempty"`

    // LoadBalancerId
    Loadbalancerid   string `json:"loadbalancerid,omitempty"`

    // Listeners
    Listeners   []Listener `json:"listeners,omitempty"`

}
