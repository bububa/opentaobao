package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelorderuserrefund 用户发起售后退款(整单/部分)
// alibaba.wdk.channel.order.userrefund
//
// 用户发起售后退款(整单/部分)
func Alibabawdkchannelorderuserrefund(clt *core.SDKClient, req *wdk.AlibabawdkchannelorderuserrefundAPIRequest, session string) (*wdk.AlibabawdkchannelorderuserrefundAPIResponse, error) {
	var resp wdk.AlibabawdkchannelorderuserrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
