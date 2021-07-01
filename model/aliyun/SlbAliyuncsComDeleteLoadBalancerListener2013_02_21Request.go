package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除 slb listener API请求
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip
*/
type SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest struct {
    model.Params
    // loadBalancerId
    _loadBalancerId   string
    // listenerPort
    _listenerPort   int64
}

// 初始化SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest对象
func NewSlbAliyuncsComDeleteLoadBalancerListener2013_02_21Request() *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest{
    return &SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetApiMethodName() string {
    return "slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21"
}

// IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
    r._loadBalancerId = _loadBalancerId
    r.Set("loadBalancerId", _loadBalancerId)
    return nil
}

// LoadBalancerId Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetLoadBalancerId() string {
    return r._loadBalancerId
}
// ListenerPort Setter
// listenerPort
func (r *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) SetListenerPort(_listenerPort int64) error {
    r._listenerPort = _listenerPort
    r.Set("listenerPort", _listenerPort)
    return nil
}

// ListenerPort Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetListenerPort() int64 {
    return r._listenerPort
}
