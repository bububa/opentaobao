package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeAgreeVtwoAPIRequest 代理商同意改签v2--增加鉴权校验 API请求
// taobao.train.agent.change.agree.vtwo
//
// 代理商同意改签接口服务
type TaobaoTrainAgentChangeAgreeVtwoAPIRequest struct {
	model.Params
	// 代理商同意改签参数
	_param *AgentAgreeChangeParam
}

// NewTaobaoTrainAgentChangeAgreeVtwoRequest 初始化TaobaoTrainAgentChangeAgreeVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeAgreeVtwoRequest() *TaobaoTrainAgentChangeAgreeVtwoAPIRequest {
	return &TaobaoTrainAgentChangeAgreeVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeAgreeVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.change.agree.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeAgreeVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 代理商同意改签参数
func (r *TaobaoTrainAgentChangeAgreeVtwoAPIRequest) SetParam(_param *AgentAgreeChangeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoTrainAgentChangeAgreeVtwoAPIRequest) GetParam() *AgentAgreeChangeParam {
	return r._param
}
