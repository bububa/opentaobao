package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请取消 API请求
alibaba.tcls.aelophy.merchant.channel.refund.cancel

翱象小程序 用户逆向申请取消
*/
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest struct {
    model.Params
    // 逆向申请取消
    _refundCancelInfo   *RefundCancelInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundCancelRequest() *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCancelInfo Setter
// 逆向申请取消
func (r *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) SetRefundCancelInfo(_refundCancelInfo *RefundCancelInfo) error {
    r._refundCancelInfo = _refundCancelInfo
    r.Set("refund_cancel_info", _refundCancelInfo)
    return nil
}

// RefundCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetRefundCancelInfo() *RefundCancelInfo {
    return r._refundCancelInfo
}
