package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配置LoadBalancer的别名。 API请求
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

// 初始化SlbAliyuncsComSetLoadBalancerName2013_02_21Request对象
func NewSlbAliyuncsComSetLoadBalancerName2013_02_21Request() *SlbAliyuncsComSetLoadBalancerName2013_02_21Request{
    return &SlbAliyuncsComSetLoadBalancerName2013_02_21Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetApiMethodName() string {
    return "slb.aliyuncs.com.SetLoadBalancerName.2013-02-21"
}

// IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerId(loadBalancerId string) error {
    r.loadBalancerId = loadBalancerId
    r.Set("loadBalancerId", loadBalancerId)
    return nil
}

// LoadBalancerId Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerId() string {
    return r.loadBalancerId
}
// LoadBalancerName Setter
// loadBalancerName
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerName(loadBalancerName string) error {
    r.loadBalancerName = loadBalancerName
    r.Set("loadBalancerName", loadBalancerName)
    return nil
}

// LoadBalancerName Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerName() string {
    return r.loadBalancerName
}
