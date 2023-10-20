package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderGetVtwoAPIRequest 代理商获取订单信息回调APIv2--增加鉴权校验 API请求
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentOrderGetVtwoAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentOrderGetVtwoRequest 初始化TaobaoTrainAgentOrderGetVtwoAPIRequest对象
func NewTaobaoTrainAgentOrderGetVtwoRequest() *TaobaoTrainAgentOrderGetVtwoAPIRequest {
	return &TaobaoTrainAgentOrderGetVtwoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentOrderGetVtwoAPIRequest) Reset() {
	r._mainOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentOrderGetVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentOrderGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoTrainAgentOrderGetVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentOrderGetVtwoRequest()
	},
}

// GetTaobaoTrainAgentOrderGetVtwoRequest 从 sync.Pool 获取 TaobaoTrainAgentOrderGetVtwoAPIRequest
func GetTaobaoTrainAgentOrderGetVtwoAPIRequest() *TaobaoTrainAgentOrderGetVtwoAPIRequest {
	return poolTaobaoTrainAgentOrderGetVtwoAPIRequest.Get().(*TaobaoTrainAgentOrderGetVtwoAPIRequest)
}

// ReleaseTaobaoTrainAgentOrderGetVtwoAPIRequest 将 TaobaoTrainAgentOrderGetVtwoAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentOrderGetVtwoAPIRequest(v *TaobaoTrainAgentOrderGetVtwoAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentOrderGetVtwoAPIRequest.Put(v)
}
