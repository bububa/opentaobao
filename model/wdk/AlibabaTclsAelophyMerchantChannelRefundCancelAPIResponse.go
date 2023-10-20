package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse 翱象商家自有渠道 逆向单申请取消 API返回值
// alibaba.tcls.aelophy.merchant.channel.refund.cancel
//
// 翱象小程序 用户逆向申请取消
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel is 翱象商家自有渠道 逆向单申请取消 成功返回结果
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse
func GetAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse() *AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse 将 AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse(v *AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse.Put(v)
}
