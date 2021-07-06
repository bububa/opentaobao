package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundAgreeAPIRequest 同意退票 API请求
// taobao.alitrip.ie.agent.refund.agree
//
// 卖家同意买家退票申请
type TaobaoAlitripIeAgentRefundAgreeAPIRequest struct {
	model.Params
	// 回复信息
	_agentAnswer string
	// 退款金额
	_refundMoney int64
	// 申请单id
	_applyId int64
	// 订单id
	_orderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoAlitripIeAgentRefundAgreeRequest 初始化TaobaoAlitripIeAgentRefundAgreeAPIRequest对象
func NewTaobaoAlitripIeAgentRefundAgreeRequest() *TaobaoAlitripIeAgentRefundAgreeAPIRequest {
	return &TaobaoAlitripIeAgentRefundAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentAnswer is AgentAnswer Setter
// 回复信息
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetAgentAnswer(_agentAnswer string) error {
	r._agentAnswer = _agentAnswer
	r.Set("agent_answer", _agentAnswer)
	return nil
}

// GetAgentAnswer AgentAnswer Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetAgentAnswer() string {
	return r._agentAnswer
}

// SetRefundMoney is RefundMoney Setter
// 退款金额
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetRefundMoney(_refundMoney int64) error {
	r._refundMoney = _refundMoney
	r.Set("refund_money", _refundMoney)
	return nil
}

// GetRefundMoney RefundMoney Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetRefundMoney() int64 {
	return r._refundMoney
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetAgentId() int64 {
	return r._agentId
}
