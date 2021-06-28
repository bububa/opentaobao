package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 交易订单取消 APIResponse
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单
*/
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyMerchantChannelOrderCancelResponse
}

type AlibabaTclsAelophyMerchantChannelOrderCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaTclsAelophyMerchantChannelOrderCancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
