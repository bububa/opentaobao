package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeRefuseVtwoAPIRequest 代理商拒绝改签v2--增加鉴权校验 API请求
// taobao.train.agent.change.refuse.vtwo
//
// 代理商拒绝火车票改签服务
type TaobaoTrainAgentChangeRefuseVtwoAPIRequest struct {
	model.Params
	// 代理商拒绝改签参数
	_param *AgentRefuseChangeParam
}

// NewTaobaoTrainAgentChangeRefuseVtwoRequest 初始化TaobaoTrainAgentChangeRefuseVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeRefuseVtwoRequest() *TaobaoTrainAgentChangeRefuseVtwoAPIRequest {
	return &TaobaoTrainAgentChangeRefuseVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeRefuseVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.change.refuse.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeRefuseVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 代理商拒绝改签参数
func (r *TaobaoTrainAgentChangeRefuseVtwoAPIRequest) SetParam(_param *AgentRefuseChangeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentChangeRefuseVtwoAPIRequest) GetParam() *AgentRefuseChangeParam {
	return r._param
}
