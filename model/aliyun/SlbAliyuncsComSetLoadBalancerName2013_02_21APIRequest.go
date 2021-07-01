package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest
配置LoadBalancer的别名。 API请求
slb.aliyuncs.com.SetLoadBalancerName.2013-02-21

配置LoadBalancer的别名。 */
type SlbAliyuncsComSetLoadBalancerName2013_02_21APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// loadBalancerName
	_loadBalancerName string
}

// New
