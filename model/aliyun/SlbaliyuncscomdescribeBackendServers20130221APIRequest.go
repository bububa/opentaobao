package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComDescribeBackendServers20130221APIRequest 后端服务器健康检查 API请求
// slb.aliyuncs.com.DescribeBackendServers.2013-02-21
//
// 后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
type SlbAliyuncsComDescribeBackendServers20130221APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// listenerPort
	_listenerPort int64
}

// NewSlbAliyuncsComDescribeBackendServers20130221Request 初始化SlbAliyuncsComDescribeBackendServers20130221APIRequest对象
func NewSlbAliyuncsComDescribeBackendServers20130221Request() *SlbAliyuncsComDescribeBackendServers20130221APIRequest {
	return &SlbAliyuncsComDescribeBackendServers20130221APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComDescribeBackendServers20130221APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.DescribeBackendServers.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComDescribeBackendServers20130221APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbAliyuncsComDescribeBackendServers20130221APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComDescribeBackendServers20130221APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbAliyuncsComDescribeBackendServers20130221APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetListenerPort is ListenerPort Setter
// listenerPort
func (r *SlbAliyuncsComDescribeBackendServers20130221APIRequest) SetListenerPort(_listenerPort int64) error {
	r._listenerPort = _listenerPort
	r.Set("listenerPort", _listenerPort)
	return nil
}

// GetListenerPort ListenerPort Getter
func (r SlbAliyuncsComDescribeBackendServers20130221APIRequest) GetListenerPort() int64 {
	return r._listenerPort
}
