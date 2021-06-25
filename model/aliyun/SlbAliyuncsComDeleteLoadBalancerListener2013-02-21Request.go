package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除 slb listener APIRequest
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip
*/
type SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request struct {
    model.Params

    // loadBalancerId
    loadBalancerId   string 

    // listenerPort
    listenerPort   int64 

}

func NewSlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request() *SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request{
    return &SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request{
        Params: model.NewParams(),
    }
}

func (r SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) GetApiMethodName() string {
    return "slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21"
}

func (r SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) SetLoadBalancerId(loadBalancerId string) error {
    r.loadBalancerId = loadBalancerId
    r.Set("loadBalancerId", loadBalancerId)
    return nil
}

func (r SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) GetLoadBalancerId() string {
    return r.loadBalancerId
}

func (r *SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) SetListenerPort(listenerPort int64) error {
    r.listenerPort = listenerPort
    r.Set("listenerPort", listenerPort)
    return nil
}

func (r SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request) GetListenerPort() int64 {
    return r.listenerPort
}

