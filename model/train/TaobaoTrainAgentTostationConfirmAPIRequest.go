package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTostationConfirmAPIRequest 线下票确认送票至车站服务 API请求
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
type TaobaoTrainAgentTostationConfirmAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoTrainAgentTostationConfirmRequest 初始化TaobaoTrainAgentTostationConfirmAPIRequest对象
func NewTaobaoTrainAgentTostationConfirmRequest() *TaobaoTrainAgentTostationConfirmAPIRequest {
	return &TaobaoTrainAgentTostationConfirmAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentTostationConfirmAPIRequest) Reset() {
	r._mainOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentTostationConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.tostation.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentTostationConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentTostationConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentTostationConfirmAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentTostationConfirmAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentTostationConfirmAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentTostationConfirmAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolTaobaoTrainAgentTostationConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentTostationConfirmRequest()
	},
}

// GetTaobaoTrainAgentTostationConfirmRequest 从 sync.Pool 获取 TaobaoTrainAgentTostationConfirmAPIRequest
func GetTaobaoTrainAgentTostationConfirmAPIRequest() *TaobaoTrainAgentTostationConfirmAPIRequest {
	return poolTaobaoTrainAgentTostationConfirmAPIRequest.Get().(*TaobaoTrainAgentTostationConfirmAPIRequest)
}

// ReleaseTaobaoTrainAgentTostationConfirmAPIRequest 将 TaobaoTrainAgentTostationConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentTostationConfirmAPIRequest(v *TaobaoTrainAgentTostationConfirmAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentTostationConfirmAPIRequest.Put(v)
}
