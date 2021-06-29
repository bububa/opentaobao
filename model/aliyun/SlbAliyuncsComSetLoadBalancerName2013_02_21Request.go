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
type SlbAliyuncsComSetLoadBalancerName2013_02_21Request struct {
    model.Params

    // loadBalancerId
    loadBalancerId   string 

    // loadBalancerName
    loadBalancerName   string 

}

func NewSlbAliyuncsComSetLoadBalancerName2013_02_21Request() *SlbAliyuncsComSetLoadBalancerName2013_02_21Request{
    return &SlbAliyuncsComSetLoadBalancerName2013_02_21Request{
        Params: model.NewParams(),
    }
}

func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetApiMethodName() string {
    return "slb.aliyuncs.com.SetLoadBalancerName.2013-02-21"
}

func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerId(loadBalancerId string) error {
    r.loadBalancerId = loadBalancerId
    r.Set("loadBalancerId", loadBalancerId)
    return nil
}

func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerId() string {
    return r.loadBalancerId
}

func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerName(loadBalancerName string) error {
    r.loadBalancerName = loadBalancerName
    r.Set("loadBalancerName", loadBalancerName)
    return nil
}

func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerName() string {
    return r.loadBalancerName
}

