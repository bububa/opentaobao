package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// SlbaliyuncscomdeleteLoadBalancerListener20130221 删除 slb listener
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
func SlbaliyuncscomdeleteLoadBalancerListener20130221(clt *core.SDKClient, req *aliyun.SlbaliyuncscomdeleteLoadBalancerListener20130221APIRequest, session string) (*aliyun.SlbaliyuncscomdeleteLoadBalancerListener20130221APIResponse, error) {
	var resp aliyun.SlbaliyuncscomdeleteLoadBalancerListener20130221APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
