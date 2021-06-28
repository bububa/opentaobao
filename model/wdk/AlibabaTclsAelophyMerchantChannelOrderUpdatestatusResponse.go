package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单状态更新 APIResponse
alibaba.tcls.aelophy.merchant.channel.order.updatestatus

订单状态变更
*/
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusResponse `json:"alibaba_tcls_aelophy_merchant_channel_order_updatestatus_response,omitempty"` 
    AlibabaTclsAelophyMerchantChannelOrderUpdatestatusResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult  *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult `json:"alibaba_tcls_aelophy_merchant_channel_order_updatestatus_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusResponse struct {

    // 返回结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusApiResult `json:"api_result,omitempty"`

}
