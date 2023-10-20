package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest 删除 slb listener API请求
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
type SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// listenerPort
	_listenerPort int64
}

// NewSlbaliyuncscomdeleteLoadBalancerListener20130221Request 初始化SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest对象
func NewSlbaliyuncscomdeleteLoadBalancerListener20130221Request() *SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest {
	return &SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetListenerPort is ListenerPort Setter
// listenerPort
func (r *SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) SetListenerPort(_listenerPort int64) error {
	r._listenerPort = _listenerPort
	r.Set("listenerPort", _listenerPort)
	return nil
}

// GetListenerPort ListenerPort Getter
func (r SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest) GetListenerPort() int64 {
	return r._listenerPort
}
