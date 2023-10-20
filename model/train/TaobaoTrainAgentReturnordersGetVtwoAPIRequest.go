package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnordersGetVtwoAPIRequest 获取待退票的订单v2--增加鉴权校验 API请求
// taobao.train.agent.returnorders.get.vtwo
//
// 代理商用来获取待退票的订单列表及数量，防止代理商掉单。
type TaobaoTrainAgentReturnordersGetVtwoAPIRequest struct {
	model.Params
	// 卖家ID
	_agentId int64
	// 0 线上退票 1线下退票
	_offline int64
}

// NewTaobaoTrainAgentReturnordersGetVtwoRequest 初始化TaobaoTrainAgentReturnordersGetVtwoAPIRequest对象
func NewTaobaoTrainAgentReturnordersGetVtwoRequest() *TaobaoTrainAgentReturnordersGetVtwoAPIRequest {
	return &TaobaoTrainAgentReturnordersGetVtwoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) Reset() {
	r._agentId = 0
	r._offline = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnorders.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 卖家ID
func (r *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetOffline is Offline Setter
// 0 线上退票 1线下退票
func (r *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) SetOffline(_offline int64) error {
	r._offline = _offline
	r.Set("offline", _offline)
	return nil
}

// GetOffline Offline Getter
func (r TaobaoTrainAgentReturnordersGetVtwoAPIRequest) GetOffline() int64 {
	return r._offline
}

var poolTaobaoTrainAgentReturnordersGetVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentReturnordersGetVtwoRequest()
	},
}

// GetTaobaoTrainAgentReturnordersGetVtwoRequest 从 sync.Pool 获取 TaobaoTrainAgentReturnordersGetVtwoAPIRequest
func GetTaobaoTrainAgentReturnordersGetVtwoAPIRequest() *TaobaoTrainAgentReturnordersGetVtwoAPIRequest {
	return poolTaobaoTrainAgentReturnordersGetVtwoAPIRequest.Get().(*TaobaoTrainAgentReturnordersGetVtwoAPIRequest)
}

// ReleaseTaobaoTrainAgentReturnordersGetVtwoAPIRequest 将 TaobaoTrainAgentReturnordersGetVtwoAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentReturnordersGetVtwoAPIRequest(v *TaobaoTrainAgentReturnordersGetVtwoAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentReturnordersGetVtwoAPIRequest.Put(v)
}
