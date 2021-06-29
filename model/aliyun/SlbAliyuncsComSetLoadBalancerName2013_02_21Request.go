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
    _loadBalancerId   string
    // loadBalancerName
    _loadBalancerName   string
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
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerId(_loadBalancerId string) error {
    r._loadBalancerId = _loadBalancerId
    r.Set("loadBalancerId", _loadBalancerId)
    return nil
}

// LoadBalancerId Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerId() string {
    return r._loadBalancerId
}
// LoadBalancerName Setter
// loadBalancerName
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21Request) SetLoadBalancerName(_loadBalancerName string) error {
    r._loadBalancerName = _loadBalancerName
    r.Set("loadBalancerName", _loadBalancerName)
    return nil
}

// LoadBalancerName Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21Request) GetLoadBalancerName() string {
    return r._loadBalancerName
}
