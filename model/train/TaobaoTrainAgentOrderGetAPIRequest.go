package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderGetAPIRequest 代理商获取订单信息回调API API请求
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentOrderGetAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentOrderGetRequest 初始化TaobaoTrainAgentOrderGetAPIRequest对象
func NewTaobaoTrainAgentOrderGetRequest() *TaobaoTrainAgentOrderGetAPIRequest {
	return &TaobaoTrainAgentOrderGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderGetAPIRequest) Reset() {
	r._mainOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentOrderGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentOrderGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentOrderGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoTrainAgentOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderGetRequest()
	},
}

// GetTaobaoTrainAgentOrderGetRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderGetAPIRequest
func GetTaobaoTrainAgentOrderGetAPIRequest() *TaobaoTrainAgentOrderGetAPIRequest {
	return poolTaobaoTrainAgentOrderGetAPIRequest.Get().(*TaobaoTrainAgentOrderGetAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderGetAPIRequest 将 TaobaoTrainAgentOrderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderGetAPIRequest(v *TaobaoTrainAgentOrderGetAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderGetAPIRequest.Put(v)
}
