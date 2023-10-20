package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// SlbAliyuncsComSetLoadBalancerName20130221 配置LoadBalancer的别名。
// slb.aliyuncs.com.SetLoadBalancerName.2013-02-21
//
// 配置LoadBalancer的别名。
func SlbAliyuncsComSetLoadBalancerName20130221(clt *core.SDKClient, req *aliyun.SlbAliyuncsComSetLoadBalancerName20130221APIRequest, session string) (*aliyun.SlbAliyuncsComSetLoadBalancerName20130221APIResponse, error) {
	var resp aliyun.SlbAliyuncsComSetLoadBalancerName20130221APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
