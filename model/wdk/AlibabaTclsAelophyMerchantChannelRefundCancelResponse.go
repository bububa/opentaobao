package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请取消 APIResponse
alibaba.tcls.aelophy.merchant.channel.refund.cancel

翱象小程序 用户逆向申请取消
*/
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_merchant_channel_refund_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundCancelApiResult `json:"api_result,omitempty" xml:"