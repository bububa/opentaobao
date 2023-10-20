package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorprefundinterceptAPIRequest 卖家发起拦截 API请求
// taobao.rp.refund.intercept
//
// 卖家发起拦截
type TaobaorprefundinterceptAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退款版本号
	_refundVersion int64
}

// NewTaobaorprefundinterceptRequest 初始化TaobaorprefundinterceptAPIRequest对象
func NewTaobaorprefundinterceptRequest() *TaobaorprefundinterceptAPIRequest {
	return &TaobaorprefundinterceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorprefundinterceptAPIRequest) GetApiMethodName() string {
	return "taobao.rp.refund.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorprefundinterceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorprefundinterceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorprefundinterceptAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorprefundinterceptAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaorprefundinterceptAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaorprefundinterceptAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}
