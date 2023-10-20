package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagenttostationreceiveAPIRequest 线下票送票至车站代理商确认用户已取票服务 API请求
// taobao.train.agent.tostation.receive
//
// 送票至车站的订单，代理商确认用户已取票
type TaobaotrainagenttostationreceiveAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagenttostationreceiveRequest 初始化TaobaotrainagenttostationreceiveAPIRequest对象
func NewTaobaotrainagenttostationreceiveRequest() *TaobaotrainagenttostationreceiveAPIRequest {
	return &TaobaotrainagenttostationreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagenttostationreceiveAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.tostation.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagenttostationreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagenttostationreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaotrainagenttostationreceiveAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagenttostationreceiveAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagenttostationreceiveAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagenttostationreceiveAPIRequest) GetAgentId() int64 {
	return r._agentId
}
