package wdk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaWdkMerchantRoutingRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantRoutingRegisterAPIResponseModel).Reset()
}

// AlibabaWdkMerchantRoutingRegisterAPIResponseModel is 商家注册更新路由信息 成功返回结果
type AlibabaWdkMerchantRoutingRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_routing_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaWdkMerchantRoutingRegisterApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantRoutingRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMerchantRoutingRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantRoutingRegisterAPIResponse)
	},
}

// GetAlibabaWdkMerchantRoutingRegisterAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantRoutingRegisterAPIResponse
func GetAlibabaWdkMerchantRoutingRegisterAPIResponse() *AlibabaWdkMerchantRoutingRegisterAPIResponse {
	return poolAlibabaWdkMerchantRoutingRegisterAPIResponse.Get().(*AlibabaWdkMerchantRoutingRegisterAPIResponse)
}

// ReleaseAlibabaWdkMerchantRoutingRegisterAPIResponse 将 AlibabaWdkMerchantRoutingRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantRoutingRegisterAPIResponse(v *AlibabaWdkMerchantRoutingRegisterAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantRoutingRegisterAPIResponse.Put(v)
}
