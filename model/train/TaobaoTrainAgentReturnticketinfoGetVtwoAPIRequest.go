package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentreturnticketinfogetvtwoAPIRequest 代理商获取退票详情回调 API请求
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
type TaobaotrainagentreturnticketinfogetvtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 淘宝子订单号
	_subOrderId int64
	// 淘宝主订单号
	_mainOrderId int64
}

// NewTaobaotrainagentreturnticketinfogetvtwoRequest 初始化TaobaotrainagentreturnticketinfogetvtwoAPIRequest对象
func NewTaobaotrainagentreturnticketinfogetvtwoRequest() *TaobaotrainagentreturnticketinfogetvtwoAPIRequest {
	return &TaobaotrainagentreturnticketinfogetvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnticketinfo.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentreturnticketinfogetvtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetSubOrderId is SubOrderId Setter
// 淘宝子订单号
func (r *TaobaotrainagentreturnticketinfogetvtwoAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单号
func (r *TaobaotrainagentreturnticketinfogetvtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagentreturnticketinfogetvtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
