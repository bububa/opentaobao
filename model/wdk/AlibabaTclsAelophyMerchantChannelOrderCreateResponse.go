package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单创建 APIResponse
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建
*/
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_merchant_channel_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult `json:"api_result,omitempty" xml:"