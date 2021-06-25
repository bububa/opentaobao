package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商同意改签v2--增加鉴权校验 APIRequest
taobao.train.agent.change.agree.vtwo

代理商同意改签接口服务
*/
type TaobaoTrainAgentChangeAgreeVtwoRequest struct {
    model.Params

    // 代理商同意改签参数
    param   *AgentAgreeChangeParam 

}

func NewTaobaoTrainAgentChangeAgreeVtwoRequest() *TaobaoTrainAgentChangeAgreeVtwoRequest{
    return &TaobaoTrainAgentChangeAgreeVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.agree.vtwo"
}

func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentChangeAgreeVtwoRequest) SetParam(param *AgentAgreeChangeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoTrainAgentChangeAgreeVtwoRequest) GetParam() *AgentAgreeChangeParam {
    return r.param
}

