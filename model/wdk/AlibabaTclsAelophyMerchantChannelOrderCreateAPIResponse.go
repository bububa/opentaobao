package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse
翱象商家自有渠道 订单创建 API返回值
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建 */
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponseModel
}

// AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponseModel is 翱象商家自有渠道 订单创建 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
