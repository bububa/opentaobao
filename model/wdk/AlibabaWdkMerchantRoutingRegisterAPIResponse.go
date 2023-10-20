package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantroutingregisterAPIResponse 商家注册更新路由信息 API返回值
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
type AlibabawdkmerchantroutingregisterAPIResponse struct {
	model.CommonResponse
	AlibabawdkmerchantroutingregisterAPIResponseModel
}

// AlibabawdkmerchantroutingregisterAPIResponseModel is 商家注册更新路由信息 成功返回结果
type AlibabawdkmerchantroutingregisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_routing_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabawdkmerchantroutingregisterApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
