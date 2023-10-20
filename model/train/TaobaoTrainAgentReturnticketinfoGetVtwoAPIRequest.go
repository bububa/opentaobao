package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest 代理商获取退票详情回调 API请求
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
type TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 淘宝子订单号
	_subOrderId int64
	// 淘宝主订单号
	_mainOrderId int64
}

// NewTaobaoTrainAgentReturnticketinfoGetVtwoRequest 初始化TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest对象
func NewTaobaoTrainAgentReturnticketinfoGetVtwoRequest() *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest {
	return &TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) Reset() {
	r._agentId = 0
	r._subOrderId = 0
	r._mainOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnticketinfo.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetSubOrderId is SubOrderId Setter
// 淘宝子订单号
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单号
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

var poolTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentReturnticketinfoGetVtwoRequest()
	},
}

// GetTaobaoTrainAgentReturnticketinfoGetVtwoRequest 从 sync.Pool 获取 TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest
func GetTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest() *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest {
	return poolTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest.Get().(*TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest)
}

// ReleaseTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest 将 TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest(v *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest.Put(v)
}
