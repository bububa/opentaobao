package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse 前置校验商品是否可下单作业 API返回值
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
type AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel is 前置校验商品是否可下单作业 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_precheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse
func GetAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse() *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse 将 AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse(v *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse.Put(v)
}
