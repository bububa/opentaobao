package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantRoutingRegister 商家注册更新路由信息
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
func AlibabaWdkMerchantRoutingRegister(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantRoutingRegisterAPIRequest, resp *wdk.AlibabaWdkMerchantRoutingRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
