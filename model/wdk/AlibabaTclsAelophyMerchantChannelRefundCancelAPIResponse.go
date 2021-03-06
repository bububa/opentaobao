package wdk

import (
	"encoding/xml"

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

// AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel is 翱象商家自有渠道 逆向单申请取消 成功返回结果
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
