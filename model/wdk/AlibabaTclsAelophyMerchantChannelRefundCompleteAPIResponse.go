package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse 翱象商家自有渠道 逆向单完成 API返回值
// alibaba.tcls.aelophy.merchant.channel.refund.complete
//
// 翱象小程序 退款完成
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponseModel is 翱象商家自有渠道 逆向单完成 成功返回结果
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse
func GetAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse() *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse 将 AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse(v *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse.Put(v)
}
