package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单完成 API请求
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成
*/
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest struct {
    model.Params
    // 请求对象
    _refundCompleteInfo   *RefundCompleteInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundCompleteRequest() *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCompleteInfo Setter
// 请求对象
func (r *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) SetRefundCompleteInfo(_refundCompleteInfo *RefundCompleteInfo) error {
    r._refundCompleteInfo = _refundCompleteInfo
    r.Set("refund_complete_info", _refundCompleteInfo)
    return nil
}

// RefundCompleteInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetRefundCompleteInfo() *RefundCompleteInfo {
    return r._refundCompleteInfo
}
