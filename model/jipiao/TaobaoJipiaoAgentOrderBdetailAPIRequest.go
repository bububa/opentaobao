package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojipiaoagentorderbdetailAPIRequest 【机票代理商订单】采购订单详情 API请求
// taobao.jipiao.agent.order.bdetail
//
// 根据淘宝系统订单号获取订单详情信息
type TaobaojipiaoagentorderbdetailAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewTaobaojipiaoagentorderbdetailRequest 初始化TaobaojipiaoagentorderbdetailAPIRequest对象
func NewTaobaojipiaoagentorderbdetailRequest() *TaobaojipiaoagentorderbdetailAPIRequest {
	return &TaobaojipiaoagentorderbdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojipiaoagentorderbdetailAPIRequest) GetApiMethodName() string {
	return "taobao.jipiao.agent.order.bdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojipiaoagentorderbdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojipiaoagentorderbdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaojipiaoagentorderbdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaojipiaoagentorderbdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}
