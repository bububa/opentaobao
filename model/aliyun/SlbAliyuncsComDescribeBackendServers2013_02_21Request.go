package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端服务器健康检查 APIRequest
slb.aliyuncs.com.DescribeBackendServers.2013-02-21

后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
*/
type SlbAliyuncsComDescribeBackendServers2013_02_21Request struct {
    model.Params

    // loadBalancerId
    loadBalancerId   string 

    // listenerPort
    listenerPort   int64 

}

func NewSlbAliyuncsComDescribeBackendServers2013_02_21Request() *SlbAliyuncsComDescribeBackendServers2013_02_21Request{
    return &SlbAliyuncsComDescribeBackendServers2013_02_21Request{
        Params: model.NewParams(),
    }
}

func (r SlbAliyuncsComDescribeBackendServers2013_02_21Request) GetApiMethodName() string {
    return "slb.aliyuncs.com.DescribeBackendServers.2013-02-21"
}

func (r SlbAliyuncsComDescribeBackendServers2013_02_21Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *SlbAliyuncsComDescribeBackendServers2013_02_21Request) SetLoadBalancerId(loadBalancerId string) error {
    r.loadBalancerId = loadBalancerId
    r.Set("loadBalancerId", loadBalancerId)
    return nil
}

func (r SlbAliyuncsComDescribeBackendServers2013_02_21Request) GetLoadBalancerId() string {
    return r.loadBalancerId
}

func (r *SlbAliyuncsComDescribeBackendServers2013_02_21Request) SetListenerPort(listenerPort int64) error {
    r.listenerPort = listenerPort
    r.Set("listenerPort", listenerPort)
    return nil
}

func (r SlbAliyuncsComDescribeBackendServers2013_02_21Request) GetListenerPort() int64 {
    return r.listenerPort
}

