package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantRoutingRegisterAPIResponse 商家注册更新路由信息 API返回值
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
type AlibabaWdkMerchantRoutingRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantRoutingRegisterAPIResponseModel
}

// AlibabaWdkMerchantRoutingRegisterAPIResponseModel is 商家注册更新路由信息 成功返回结果
type AlibabaWdkMerchantRoutingRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_routing_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaWdkMerchantRoutingRegisterApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
