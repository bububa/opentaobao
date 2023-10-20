package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundDisagreeAPIRequest saas 售后逆向 商户拒绝用户逆向申请 API请求
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
type AlibabaTclsAelophyRefundDisagreeAPIRequest struct {
	model.Params
	// 退款单ID
	_refundId string
	// 拒绝原因
	_rejectReason string
	// 渠道来源
	_orderFrom int64
}

// NewAlibabaTclsAelophyRefundDisagreeRequest 初始化AlibabaTclsAelophyRefundDisagreeAPIRequest对象
func NewAlibabaTclsAelophyRefundDisagreeRequest() *AlibabaTclsAelophyRefundDisagreeAPIRequest {
	return &AlibabaTclsAelophyRefundDisagreeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) Reset() {
	r._refundId = ""
	r._rejectReason = ""
	r._orderFrom = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.disagree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetRejectReason is RejectReason Setter
// 拒绝原因
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) SetRejectReason(_rejectReason string) error {
	r._rejectReason = _rejectReason
	r.Set("reject_reason", _rejectReason)
	return nil
}

// GetRejectReason RejectReason Getter
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetRejectReason() string {
	return r._rejectReason
}

// SetOrderFrom is OrderFrom Setter
// 渠道来源
func (r *AlibabaTclsAelophyRefundDisagreeAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabaTclsAelophyRefundDisagreeAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}

var poolAlibabaTclsAelophyRefundDisagreeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyRefundDisagreeRequest()
	},
}

// GetAlibabaTclsAelophyRefundDisagreeRequest 从 sync.Pool 获取 AlibabaTclsAelophyRefundDisagreeAPIRequest
func GetAlibabaTclsAelophyRefundDisagreeAPIRequest() *AlibabaTclsAelophyRefundDisagreeAPIRequest {
	return poolAlibabaTclsAelophyRefundDisagreeAPIRequest.Get().(*AlibabaTclsAelophyRefundDisagreeAPIRequest)
}

// ReleaseAlibabaTclsAelophyRefundDisagreeAPIRequest 将 AlibabaTclsAelophyRefundDisagreeAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyRefundDisagreeAPIRequest(v *AlibabaTclsAelophyRefundDisagreeAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyRefundDisagreeAPIRequest.Put(v)
}
