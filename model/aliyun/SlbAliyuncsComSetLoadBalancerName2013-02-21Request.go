package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配置LoadBalancer的别名。 APIRequest
slb.aliyuncs.com.SetLoadBalancerName.2013-02-21

配置LoadBalancer的别名。
*/
type SlbAliyuncsComSetLoadBalancerName2013-02-21Request struct {
    model.Params

    // loadBalancerId
    loadBalancerId   string 

    // loadBalancerName
    loadBalancerName   string 

}

func NewSlbAliyuncsComSetLoadBalancerName2013-02-21Request() *SlbAliyuncsComSetLoadBalancerName2013-02-21Request{
    return &SlbAliyuncsComSetLoadBalancerName2013-02-21Request{
        Params: model.NewParams(),
    }
}

func (r SlbAliyuncsComSetLoadBalancerName2013-02-21Request) GetApiMethodName() string {
    return "slb.aliyuncs.com.SetLoadBalancerName.2013-02-21"
}

func (r SlbAliyuncsComSetLoadBalancerName2013-02-21Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *SlbAliyuncsComSetLoadBalancerName2013-02-21Request) SetLoadBalancerId(loadBalancerId string) error {
    r.loadBalancerId = loadBalancerId
    r.Set("loadBalancerId", loadBalancerId)
    return nil
}

func (r SlbAliyuncsComSetLoadBalancerName2013-02-21Request) GetLoadBalancerId() string {
    return r.loadBalancerId
}

func (r *SlbAliyuncsComSetLoadBalancerName2013-02-21Request) SetLoadBalancerName(loadBalancerName string) error {
    r.loadBalancerName = loadBalancerName
    r.Set("loadBalancerName", loadBalancerName)
    return nil
}

func (r SlbAliyuncsComSetLoadBalancerName2013-02-21Request) GetLoadBalancerName() string {
    return r.loadBalancerName
}

