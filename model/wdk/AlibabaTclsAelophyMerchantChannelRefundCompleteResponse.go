package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单完成 APIResponse
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_merchant_channel_refund_complete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult `json:"api_result,omitempty" xml:"