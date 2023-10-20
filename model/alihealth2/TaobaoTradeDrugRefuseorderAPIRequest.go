package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedrugrefuseorderAPIRequest 阿里健康020拒单 API请求
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
type TaobaotradedrugrefuseorderAPIRequest struct {
	model.Params
	// 拒单原因
	_refuseReason string
	// 订单ID
	_orderId int64
}

// NewTaobaotradedrugrefuseorderRequest 初始化TaobaotradedrugrefuseorderAPIRequest对象
func NewTaobaotradedrugrefuseorderRequest() *TaobaotradedrugrefuseorderAPIRequest {
	return &TaobaotradedrugrefuseorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradedrugrefuseorderAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.refuseorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradedrugrefuseorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradedrugrefuseorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefuseReason is RefuseReason Setter
// 拒单原因
func (r *TaobaotradedrugrefuseorderAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r TaobaotradedrugrefuseorderAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaotradedrugrefuseorderAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaotradedrugrefuseorderAPIRequest) GetOrderId() int64 {
	return r._orderId
}
