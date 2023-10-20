package ieagency

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundGetAPIRequest 获取退票申请详情 API请求
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
type TaobaoAlitripIeAgentRefundGetAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoAlitripIeAgentRefundGetRequest 初始化TaobaoAlitripIeAgentRefundGetAPIRequest对象
func NewTaobaoAlitripIeAgentRefundGetRequest() *TaobaoAlitripIeAgentRefundGetAPIRequest {
	return &TaobaoAlitripIeAgentRefundGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentRefundGetAPIRequest) Reset() {
	r._applyId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundGetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoAlitripIeAgentRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentRefundGetRequest()
	},
}

// GetTaobaoAlitripIeAgentRefundGetRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundGetAPIRequest
func GetTaobaoAlitripIeAgentRefundGetAPIRequest() *TaobaoAlitripIeAgentRefundGetAPIRequest {
	return poolTaobaoAlitripIeAgentRefundGetAPIRequest.Get().(*TaobaoAlitripIeAgentRefundGetAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentRefundGetAPIRequest 将 TaobaoAlitripIeAgentRefundGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundGetAPIRequest(v *TaobaoAlitripIeAgentRefundGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundGetAPIRequest.Put(v)
}
