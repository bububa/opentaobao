package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest 删除 slb listener API请求
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
type SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// listenerPort
	_listenerPort int64
}

// NewSlbAliyuncsComDeleteLoadBalancerListener20130221Request 初始化SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest对象
func NewSlbAliyuncsComDeleteLoadBalancerListener20130221Request() *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest {
	return &SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) Reset() {
	r._loadBalancerId = ""
	r._listenerPort = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetListenerPort is ListenerPort Setter
// listenerPort
func (r *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) SetListenerPort(_listenerPort int64) error {
	r._listenerPort = _listenerPort
	r.Set("listenerPort", _listenerPort)
	return nil
}

// GetListenerPort ListenerPort Getter
func (r SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) GetListenerPort() int64 {
	return r._listenerPort
}

var poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest = sync.Pool{
	New: func() any {
		return NewSlbAliyuncsComDeleteLoadBalancerListener20130221Request()
	},
}

// GetSlbAliyuncsComDeleteLoadBalancerListener20130221Request 从 sync.Pool 获取 SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest
func GetSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest() *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest {
	return poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest.Get().(*SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest)
}

// ReleaseSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest 将 SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest 放入 sync.Pool
func ReleaseSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest(v *SlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest) {
	v.Reset()
	poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIRequest.Put(v)
}
