package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单创建 APIResponse
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建
*/
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyMerchantChannelOrderCreateResponse `json:"alibaba_tcls_aelophy_merchant_channel_order_create_response,omitempty"` 
    AlibabaTclsAelophyMerchantChannelOrderCreateResponse
}

/* model for simplify = false
type AlibabaTclsAelophyMerchantChannelOrderCreateResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyMerchantChannelOrderCreateApiResult  *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult `json:"alibaba_tcls_aelophy_merchant_channel_order_create_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyMerchantChannelOrderCreateResponse struct {

    // 返回结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelOrderCreateApiResult `json:"api_result,omitempty"`

}
