package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefunddisagreeAPIRequest saas 售后逆向 商户拒绝用户逆向申请 API请求
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
type AlibabatclsaelophyrefunddisagreeAPIRequest struct {
	model.Params
	// 退款单ID
	_refundId string
	// 拒绝原因
	_rejectReason string
	// 渠道来源
	_orderFrom int64
}

// NewAlibabatclsaelophyrefunddisagreeRequest 初始化AlibabatclsaelophyrefunddisagreeAPIRequest对象
func NewAlibabatclsaelophyrefunddisagreeRequest() *AlibabatclsaelophyrefunddisagreeAPIRequest {
	return &AlibabatclsaelophyrefunddisagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.disagree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *AlibabatclsaelophyrefunddisagreeAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetRejectReason is RejectReason Setter
// 拒绝原因
func (r *AlibabatclsaelophyrefunddisagreeAPIRequest) SetRejectReason(_rejectReason string) error {
	r._rejectReason = _rejectReason
	r.Set("reject_reason", _rejectReason)
	return nil
}

// GetRejectReason RejectReason Getter
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetRejectReason() string {
	return r._rejectReason
}

// SetOrderFrom is OrderFrom Setter
// 渠道来源
func (r *AlibabatclsaelophyrefunddisagreeAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabatclsaelophyrefunddisagreeAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}
