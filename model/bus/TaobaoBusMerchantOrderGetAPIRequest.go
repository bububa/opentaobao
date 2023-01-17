package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusMerchantOrderGetAPIRequest 商家侧查询订单详情 API请求
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaoBusMerchantOrderGetAPIRequest struct {
	model.Params
	// 入参
	_paramAgentQueryOrderRQ *AgentQueryOrderRq
}

// NewTaobaoBusMerchantOrderGetRequest 初始化TaobaoBusMerchantOrderGetAPIRequest对象
func NewTaobaoBusMerchantOrderGetRequest() *TaobaoBusMerchantOrderGetAPIRequest {
	return &TaobaoBusMerchantOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusMerchantOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.merchant.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusMerchantOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusMerchantOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentQueryOrderRQ is ParamAgentQueryOrderRQ Setter
// 入参
func (r *TaobaoBusMerchantOrderGetAPIRequest) SetParamAgentQueryOrderRQ(_paramAgentQueryOrderRQ *AgentQueryOrderRq) error {
	r._paramAgentQueryOrderRQ = _paramAgentQueryOrderRQ
	r.Set("param_agent_query_order_r_q", _paramAgentQueryOrderRQ)
	return nil
}

// GetParamAgentQueryOrderRQ ParamAgentQueryOrderRQ Getter
func (r TaobaoBusMerchantOrderGetAPIRequest) GetParamAgentQueryOrderRQ() *AgentQueryOrderRq {
	return r._paramAgentQueryOrderRQ
}
