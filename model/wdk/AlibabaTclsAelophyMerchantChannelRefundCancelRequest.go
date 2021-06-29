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
type AlibabaTclsAelophyMerchantChannelRefundCancelRequest struct {
    model.Params
    // 逆向申请取消
    refundCancelInfo   *RefundCancelInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelRefundCancelRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundCancelRequest() *AlibabaTclsAelophyMerchantChannelRefundCancelRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundCancelRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCancelInfo Setter
// 逆向申请取消
func (r *AlibabaTclsAelophyMerchantChannelRefundCancelRequest) SetRefundCancelInfo(refundCancelInfo *RefundCancelInfo) error {
    r.refundCancelInfo = refundCancelInfo
    r.Set("refund_cancel_info", refundCancelInfo)
    return nil
}

// RefundCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundCancelRequest) GetRefundCancelInfo() *RefundCancelInfo {
    return r.refundCancelInfo
}
