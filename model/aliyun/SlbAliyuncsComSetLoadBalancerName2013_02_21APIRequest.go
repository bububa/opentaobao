package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest 配置LoadBalancer的别名。 API请求
// slb.aliyuncs.com.SetLoadBalancerName.2013-02-21
//
// 配置LoadBalancer的别名。
type SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// loadBalancerName
	_loadBalancerName string
}

// NewSlbAliyuncsComSetLoadBalancerName2013_02_21Request 初始化SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest对象
func NewSlbAliyuncsComSetLoadBalancerName2013_02_21Request() *SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest {
	return &SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.SetLoadBalancerName.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetLoadBalancerName is LoadBalancerName Setter
// loadBalancerName
func (r *SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) SetLoadBalancerName(_loadBalancerName string) error {
	r._loadBalancerName = _loadBalancerName
	r.Set("loadBalancerName", _loadBalancerName)
	return nil
}

// GetLoadBalancerName LoadBalancerName Getter
func (r SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest) GetLoadBalancerName() string {
	return r._loadBalancerName
}
