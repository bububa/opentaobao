package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商拒绝改签v2--增加鉴权校验 APIRequest
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务
*/
type TaobaoTrainAgentChangeRefuseVtwoRequest struct {
    model.Params

    // 代理商拒绝改签参数
    param   *AgentRefuseChangeParam 

}

func NewTaobaoTrainAgentChangeRefuseVtwoRequest() *TaobaoTrainAgentChangeRefuseVtwoRequest{
    return &TaobaoTrainAgentChangeRefuseVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.refuse.vtwo"
}

func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentChangeRefuseVtwoRequest) SetParam(param *AgentRefuseChangeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetParam() *AgentRefuseChangeParam {
    return r.param
}

