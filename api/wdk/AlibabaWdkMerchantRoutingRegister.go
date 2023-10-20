package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantroutingregister 商家注册更新路由信息
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
func Alibabawdkmerchantroutingregister(clt *core.SDKClient, req *wdk.AlibabawdkmerchantroutingregisterAPIRequest, session string) (*wdk.AlibabawdkmerchantroutingregisterAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantroutingregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
