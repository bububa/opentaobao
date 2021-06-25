package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单完成 APIRequest
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
type AlibabaTclsAelophyMerchantChannelRefundCompleteRequest struct {
    model.Params

    // 请求对象
    refundCompleteInfo   *RefundCompleteInfo 

}

func NewAlibabaTclsAelophyMerchantChannelRefundCompleteRequest() *AlibabaTclsAelophyMerchantChannelRefundCompleteRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantChannelRefundCompleteRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.complete"
}

func (r AlibabaTclsAelophyMerchantChannelRefundCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantChannelRefundCompleteRequest) SetRefundCompleteInfo(refundCompleteInfo *RefundCompleteInfo) error {
    r.refundCompleteInfo = refundCompleteInfo
    r.Set("refund_complete_info", refundCompleteInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantChannelRefundCompleteRequest) GetRefundCompleteInfo() *RefundCompleteInfo {
    return r.refundCompleteInfo
}

