package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusmerchantordergetAPIRequest 商家侧查询订单详情 API请求
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaobusmerchantordergetAPIRequest struct {
	model.Params
	// 入参
	_paramAgentQueryOrderRQ *AgentQueryOrderRq
}

// NewTaobaobusmerchantordergetRequest 初始化TaobaobusmerchantordergetAPIRequest对象
func NewTaobaobusmerchantordergetRequest() *TaobaobusmerchantordergetAPIRequest {
	return &TaobaobusmerchantordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusmerchantordergetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.merchant.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusmerchantordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusmerchantordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentQueryOrderRQ is ParamAgentQueryOrderRQ Setter
// 入参
func (r *TaobaobusmerchantordergetAPIRequest) SetParamAgentQueryOrderRQ(_paramAgentQueryOrderRQ *AgentQueryOrderRq) error {
	r._paramAgentQueryOrderRQ = _paramAgentQueryOrderRQ
	r.Set("param_agent_query_order_r_q", _paramAgentQueryOrderRQ)
	return nil
}

// GetParamAgentQueryOrderRQ ParamAgentQueryOrderRQ Getter
func (r TaobaobusmerchantordergetAPIRequest) GetParamAgentQueryOrderRQ() *AgentQueryOrderRq {
	return r._paramAgentQueryOrderRQ
}
