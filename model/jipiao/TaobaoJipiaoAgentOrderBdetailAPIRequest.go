package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJipiaoAgentOrderBdetailAPIRequest
【机票代理商订单】采购订单详情 API请求
taobao.jipiao.agent.order.bdetail

根据淘宝系统订单号获取订单详情信息 */
type TaobaoJipiaoAgentOrderBdetailAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewTaobaoJipiaoAgentOrderBdetailRequest 初始化TaobaoJipiaoAgentOrderBdetailAPIRequest对象
func NewTaobaoJipiaoAgentOrderBdetailRequest() *TaobaoJipiaoAgentOrderBdetailAPIRequest {
	return &TaobaoJipiaoAgentOrderBdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetApiMethodName() string {
	return "taobao.jipiao.agent.order.bdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单号
func (r *TaobaoJipiaoAgentOrderBdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}
