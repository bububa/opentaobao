package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单完成 API返回值 
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyMerchantChannelRefundCompleteResponse
}

// 翱象商家自有渠道 逆向单完成 成功返回结果
type AlibabaTclsAelophyMerchantChannelRefundCompleteResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_complete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    ApiResult   *AlibabaTclsAelophyMerchantChannelRefundCompleteApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
