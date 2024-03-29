package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJipiaoAgentOrderBdetailAPIRequest 【机票代理商订单】采购订单详情 API请求
// taobao.jipiao.agent.order.bdetail
//
// 根据淘宝系统订单号获取订单详情信息
type TaobaoJipiaoAgentOrderBdetailAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewTaobaoJipiaoAgentOrderBdetailRequest 初始化TaobaoJipiaoAgentOrderBdetailAPIRequest对象
func NewTaobaoJipiaoAgentOrderBdetailRequest() *TaobaoJipiaoAgentOrderBdetailAPIRequest {
	return &TaobaoJipiaoAgentOrderBdetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJipiaoAgentOrderBdetailAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetApiMethodName() string {
	return "taobao.jipiao.agent.order.bdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaoJipiaoAgentOrderBdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoJipiaoAgentOrderBdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoJipiaoAgentOrderBdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJipiaoAgentOrderBdetailRequest()
	},
}

// GetTaobaoJipiaoAgentOrderBdetailRequest 从 sync.Pool 获取 TaobaoJipiaoAgentOrderBdetailAPIRequest
func GetTaobaoJipiaoAgentOrderBdetailAPIRequest() *TaobaoJipiaoAgentOrderBdetailAPIRequest {
	return poolTaobaoJipiaoAgentOrderBdetailAPIRequest.Get().(*TaobaoJipiaoAgentOrderBdetailAPIRequest)
}

// ReleaseTaobaoJipiaoAgentOrderBdetailAPIRequest 将 TaobaoJipiaoAgentOrderBdetailAPIRequest 放入 sync.Pool
func ReleaseTaobaoJipiaoAgentOrderBdetailAPIRequest(v *TaobaoJipiaoAgentOrderBdetailAPIRequest) {
	v.Reset()
	poolTaobaoJipiaoAgentOrderBdetailAPIRequest.Put(v)
}
