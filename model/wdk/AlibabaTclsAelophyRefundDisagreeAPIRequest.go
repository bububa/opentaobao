package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户拒绝用户逆向申请 API请求
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请
*/
type AlibabaTclsAelophyRefundDisagreeAPIRequest struct {
    model.Params
    // 退款单ID
    _refundId   string
    // 拒绝原因
    _rejectReason   string
}

// 初始化AlibabaTclsAelophyRefundDisagreeAPIRequest对象
func NewAlibabaTclsAelophyRefundDisagreeRequest() *AlibabaTclsAelophyRefundDisagreeAPIRequest{
    return &AlibabaTclsAelophyRefundDisagreeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.disagree"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) SetRefundId(_refundId string) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetRefundId() string {
    return r._refundId
}
// RejectReason Setter
// 拒绝原因
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) SetRejectReason(_rejectReason string) error {
    r._rejectReason = _rejectReason
    r.Set("reject_reason", _rejectReason)
    return nil
}

// RejectReason Getter
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetRejectReason() string {
    return r._rejectReason
}
