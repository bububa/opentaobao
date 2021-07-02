package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewReceiveAPIRequest 商家退票受理申请(对外) API请求
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
type TaobaoAlitripIeAgentRefundNewReceiveAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
	// 订单id
	_orderId int64
}

// NewTaobaoAlitripIeAgentRefundNewReceiveRequest 初始化TaobaoAlitripIeAgentRefundNewReceiveAPIRequest对象
func NewTaobaoAlitripIeAgentRefundNewReceiveRequest() *TaobaoAlitripIeAgentRefundNewReceiveAPIRequest {
	return &TaobaoAlitripIeAgentRefundNewReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveAPIRequest) GetOrderId() int64 {
	return r._orderId
}
