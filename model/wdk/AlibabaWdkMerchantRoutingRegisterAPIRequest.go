package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantRoutingRegisterAPIRequest
商家注册更新路由信息 API请求
alibaba.wdk.merchant.routing.register

商家注册更新路由信息 */
type AlibabaWdkMerchantRoutingRegisterAPIRequest struct {
	model.Params
	// 接口入参
	_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo
}

// New
