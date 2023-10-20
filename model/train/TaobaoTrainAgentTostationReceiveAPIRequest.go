package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTostationReceiveAPIRequest 线下票送票至车站代理商确认用户已取票服务 API请求
// taobao.train.agent.tostation.receive
//
// 送票至车站的订单，代理商确认用户已取票
type TaobaoTrainAgentTostationReceiveAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentTostationReceiveRequest 初始化TaobaoTrainAgentTostationReceiveAPIRequest对象
func NewTaobaoTrainAgentTostationReceiveRequest() *TaobaoTrainAgentTostationReceiveAPIRequest {
	return &TaobaoTrainAgentTostationReceiveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentTostationReceiveAPIRequest) Reset() {
	r._mainOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentTostationReceiveAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.tostation.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentTostationReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentTostationReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentTostationReceiveAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentTostationReceiveAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentTostationReceiveAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentTostationReceiveAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoTrainAgentTostationReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentTostationReceiveRequest()
	},
}

// GetTaobaoTrainAgentTostationReceiveRequest 从 sync.Pool 获取 TaobaoTrainAgentTostationReceiveAPIRequest
func GetTaobaoTrainAgentTostationReceiveAPIRequest() *TaobaoTrainAgentTostationReceiveAPIRequest {
	return poolTaobaoTrainAgentTostationReceiveAPIRequest.Get().(*TaobaoTrainAgentTostationReceiveAPIRequest)
}

// ReleaseTaobaoTrainAgentTostationReceiveAPIRequest 将 TaobaoTrainAgentTostationReceiveAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentTostationReceiveAPIRequest(v *TaobaoTrainAgentTostationReceiveAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentTostationReceiveAPIRequest.Put(v)
}
