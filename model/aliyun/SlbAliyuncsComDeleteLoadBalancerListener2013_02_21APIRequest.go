package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest 删除 slb listener API请求
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
type SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// listenerPort
	_listenerPort int64
}

// NewSlbAliyuncsComDeleteLoadBalancerListener2013_02_21Request 初始化SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest对象
func NewSlbAliyuncsComDeleteLoadBalancerListener2013_02_21Request() *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest {
	return &SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetListenerPort is ListenerPort Setter
// listenerPort
func (r *SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) SetListenerPort(_listenerPort int64) error {
	r._listenerPort = _listenerPort
	r.Set("listenerPort", _listenerPort)
	return nil
}

// GetListenerPort ListenerPort Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest) GetListenerPort() int64 {
	return r._listenerPort
}
