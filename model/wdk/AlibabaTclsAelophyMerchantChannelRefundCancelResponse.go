package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请取消 APIResponse
alibaba.tcls.aelophy.merchant.channel.refund.cancel

翱象小程序 用户逆向申请取消
*/
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantChannelRefundCancelResponse `json:"alibaba_tcls_aelophy_merchant_channel_refund_cancel_response,omitempty"` 
    AlibabaTclsAelophyMerchantChannelRefundCancelResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantChannelRefundCancelResponse struct {

    // 结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantChannelRefundCancelApiResult  *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult `json:"alibaba_tcls_aelophy_merchant_channel_refund_cancel_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantChannelRefundCancelResponse struct {

    // 结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult `json:"api_result,omitempty"`

}
