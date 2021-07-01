package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest
删除 slb listener API请求
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip */
type SlbAliyuncsComDeleteLoadBalancerListener2013_02_21APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// listenerPort
	_listenerPort int64
}

// New
