package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 交易订单取消 APIResponse
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单
*/
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTclsAelophyMerchantChannelOrderCancelResponse `json:"alibaba_tcls_aelophy_merchant_channel_order_cancel_response,omitempty"`
}

type AlibabaTclsAelophyMerchantChannelOrderCancelResponse struct {

    // 返回结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelOrderCancelApiResult `json:"api_result,omitempty"`

}
