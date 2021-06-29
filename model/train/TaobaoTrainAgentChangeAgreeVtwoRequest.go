package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商同意改签v2--增加鉴权校验 API请求
taobao.train.agent.change.agree.vtwo

代理商同意改签接口服务
*/
type TaobaoTrainAgentChangeAgreeVtwoRequest struct {
    model.Params
    // 代理商同意改签参数
    param   *AgentAgreeChangeParam
}

// 初始化TaobaoTrainAgentChangeAgreeVtwoRequest对象
func NewTaobaoTrainAgentChangeAgreeVtwoRequest() *TaobaoTrainAgentChangeAgreeVtwoRequest{
    return &TaobaoTrainAgentChangeAgreeVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.agree.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 代理商同意改签参数
func (r *TaobaoTrainAgentChangeAgreeVtwoRequest) SetParam(param *AgentAgreeChangeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetParam() *AgentAgreeChangeParam {
    return r.param
}
