package ieagency

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundRefuseAPIRequest 拒绝退票申请 API请求
// taobao.alitrip.ie.agent.refund.refuse
//
// 卖家拒绝退票退票申请
type TaobaoAlitripIeAgentRefundRefuseAPIRequest struct {
	model.Params
	// 代理商回复
	_agentAnswer string
	// 退票申请单id
	_applyId int64
	// 订单id
	_orderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoAlitripIeAgentRefundRefuseRequest 初始化TaobaoAlitripIeAgentRefundRefuseAPIRequest对象
func NewTaobaoAlitripIeAgentRefundRefuseRequest() *TaobaoAlitripIeAgentRefundRefuseAPIRequest {
	return &TaobaoAlitripIeAgentRefundRefuseAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentRefundRefuseAPIRequest) Reset() {
	r._agentAnswer = ""
	r._applyId = 0
	r._orderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentAnswer is AgentAnswer Setter
// 代理商回复
func (r *TaobaoAlitripIeAgentRefundRefuseAPIRequest) SetAgentAnswer(_agentAnswer string) error {
	r._agentAnswer = _agentAnswer
	r.Set("agent_answer", _agentAnswer)
	return nil
}

// GetAgentAnswer AgentAnswer Getter
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetAgentAnswer() string {
	return r._agentAnswer
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundRefuseAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundRefuseAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefuseAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefuseAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoAlitripIeAgentRefundRefuseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentRefundRefuseRequest()
	},
}

// GetTaobaoAlitripIeAgentRefundRefuseRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundRefuseAPIRequest
func GetTaobaoAlitripIeAgentRefundRefuseAPIRequest() *TaobaoAlitripIeAgentRefundRefuseAPIRequest {
	return poolTaobaoAlitripIeAgentRefundRefuseAPIRequest.Get().(*TaobaoAlitripIeAgentRefundRefuseAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentRefundRefuseAPIRequest 将 TaobaoAlitripIeAgentRefundRefuseAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundRefuseAPIRequest(v *TaobaoAlitripIeAgentRefundRefuseAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundRefuseAPIRequest.Put(v)
}
