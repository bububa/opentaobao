package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请 APIResponse
alibaba.tcls.aelophy.merchant.channel.refund.apply

翱象小程序 用户逆向单申请
*/
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantChannelRefundApplyResponse `json:"alibaba_tcls_aelophy_merchant_channel_refund_apply_response,omitempty"` 
    AlibabaTclsAelophyMerchantChannelRefundApplyResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantChannelRefundApplyResponse struct {

    // 结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantChannelRefundApplyApiResult  *AlibabaTclsAelophyMerchantChannelRefundApplyApiResult `json:"alibaba_tcls_aelophy_merchant_channel_refund_apply_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantChannelRefundApplyResponse struct {

    // 结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundApplyApiResult `json:"api_result,omitempty"`

}
