package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse 翱象商家自有渠道 逆向单申请 API返回值
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponseModel is 翱象商家自有渠道 逆向单申请 成功返回结果
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelRefundApplyApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse
func GetAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse() *AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse 将 AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse(v *AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse.Put(v)
}
