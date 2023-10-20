package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse 翱象商家自有渠道 交易订单取消 API返回值
// alibaba.tcls.aelophy.merchant.channel.order.cancel
//
// 翱象小程序用户取消订单
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponseModel is 翱象商家自有渠道 交易订单取消 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderCancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse
func GetAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse() *AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse 将 AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse(v *AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse.Put(v)
}
