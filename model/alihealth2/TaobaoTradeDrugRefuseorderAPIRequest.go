package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugRefuseorderAPIRequest 阿里健康020拒单 API请求
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
type TaobaoTradeDrugRefuseorderAPIRequest struct {
	model.Params
	// 拒单原因
	_refuseReason string
	// 订单ID
	_orderId int64
}

// NewTaobaoTradeDrugRefuseorderRequest 初始化TaobaoTradeDrugRefuseorderAPIRequest对象
func NewTaobaoTradeDrugRefuseorderRequest() *TaobaoTradeDrugRefuseorderAPIRequest {
	return &TaobaoTradeDrugRefuseorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.refuseorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefuseReason is RefuseReason Setter
// 拒单原因
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetOrderId() int64 {
	return r._orderId
}
