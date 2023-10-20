package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelordercancelAPIResponse 翱象商家自有渠道 交易订单取消 API返回值
// alibaba.tcls.aelophy.merchant.channel.order.cancel
//
// 翱象小程序用户取消订单
type AlibabatclsaelophymerchantchannelordercancelAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantchannelordercancelAPIResponseModel
}

// AlibabatclsaelophymerchantchannelordercancelAPIResponseModel is 翱象商家自有渠道 交易订单取消 成功返回结果
type AlibabatclsaelophymerchantchannelordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabatclsaelophymerchantchannelordercancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
