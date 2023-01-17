package wdk

import (
	"encoding/xml"

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

// AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel is 前置校验商品是否可下单作业 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_precheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderPrecheckApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
