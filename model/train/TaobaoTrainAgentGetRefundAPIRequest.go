package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentGetRefundAPIRequest 代理商获取订单退票信息 API请求
// taobao.train.agent.get.refund
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentGetRefundAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentGetRefundRequest 初始化TaobaoTrainAgentGetRefundAPIRequest对象
func NewTaobaoTrainAgentGetRefundRequest() *TaobaoTrainAgentGetRefundAPIRequest {
	return &TaobaoTrainAgentGetRefundAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentGetRefundAPIRequest) Reset() {
	r._mainOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentGetRefundAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.get.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentGetRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentGetRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentGetRefundAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentGetRefundAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentGetRefundAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentGetRefundAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoTrainAgentGetRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentGetRefundRequest()
	},
}

// GetTaobaoTrainAgentGetRefundRequest 从 sync.Pool 获取 TaobaoTrainAgentGetRefundAPIRequest
func GetTaobaoTrainAgentGetRefundAPIRequest() *TaobaoTrainAgentGetRefundAPIRequest {
	return poolTaobaoTrainAgentGetRefundAPIRequest.Get().(*TaobaoTrainAgentGetRefundAPIRequest)
}

// ReleaseTaobaoTrainAgentGetRefundAPIRequest 将 TaobaoTrainAgentGetRefundAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentGetRefundAPIRequest(v *TaobaoTrainAgentGetRefundAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentGetRefundAPIRequest.Put(v)
}
