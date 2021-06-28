package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单完成 APIResponse
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantChannelRefundCompleteResponse `json:"alibaba_tcls_aelophy_merchant_channel_refund_complete_response,omitempty"` 
    AlibabaTclsAelophyMerchantChannelRefundCompleteResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantChannelRefundCompleteResponse struct {

    // 结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult  *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult `json:"alibaba_tcls_aelophy_merchant_channel_refund_complete_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantChannelRefundCompleteResponse struct {

    // 结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult `json:"api_result,omitempty"`

}
