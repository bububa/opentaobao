package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttraderefundordergetAPIRequest 零售通退款订单查询 API请求
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
type AlibabalsttraderefundordergetAPIRequest struct {
	model.Params
	// 退款单id
	_refundId string
	// 主订单id
	_mainOrderId int64
}

// NewAlibabalsttraderefundordergetRequest 初始化AlibabalsttraderefundordergetAPIRequest对象
func NewAlibabalsttraderefundordergetRequest() *AlibabalsttraderefundordergetAPIRequest {
	return &AlibabalsttraderefundordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttraderefundordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.refund.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttraderefundordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttraderefundordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款单id
func (r *AlibabalsttraderefundordergetAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabalsttraderefundordergetAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabalsttraderefundordergetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabalsttraderefundordergetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
