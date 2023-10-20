package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundnewreceiveAPIRequest 商家退票受理申请(对外) API请求
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
type TaobaoalitripieagentrefundnewreceiveAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
	// 订单id
	_orderId int64
}

// NewTaobaoalitripieagentrefundnewreceiveRequest 初始化TaobaoalitripieagentrefundnewreceiveAPIRequest对象
func NewTaobaoalitripieagentrefundnewreceiveRequest() *TaobaoalitripieagentrefundnewreceiveAPIRequest {
	return &TaobaoalitripieagentrefundnewreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoalitripieagentrefundnewreceiveAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoalitripieagentrefundnewreceiveAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoalitripieagentrefundnewreceiveAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripieagentrefundnewreceiveAPIRequest) GetOrderId() int64 {
	return r._orderId
}
