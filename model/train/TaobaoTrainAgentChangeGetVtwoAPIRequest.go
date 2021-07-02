package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeGetVtwoAPIRequest 获取改签单详情v2--增加鉴权校验 API请求
// taobao.train.agent.change.get.vtwo
//
// 卖家获取待处理的改签单详情
type TaobaoTrainAgentChangeGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 申请单id
	_applyId int64
}

// NewTaobaoTrainAgentChangeGetVtwoRequest 初始化TaobaoTrainAgentChangeGetVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeGetVtwoRequest() *TaobaoTrainAgentChangeGetVtwoAPIRequest {
	return &TaobaoTrainAgentChangeGetVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.change.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeGetVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentChangeGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentChangeGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *TaobaoTrainAgentChangeGetVtwoAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoTrainAgentChangeGetVtwoAPIRequest) GetApplyId() int64 {
	return r._applyId
}
