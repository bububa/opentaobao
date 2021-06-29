package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商拒绝改签v2--增加鉴权校验 API请求
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务
*/
type TaobaoTrainAgentChangeRefuseVtwoRequest struct {
    model.Params
    // 代理商拒绝改签参数
    _param   *AgentRefuseChangeParam
}

// 初始化TaobaoTrainAgentChangeRefuseVtwoRequest对象
func NewTaobaoTrainAgentChangeRefuseVtwoRequest() *TaobaoTrainAgentChangeRefuseVtwoRequest{
    return &TaobaoTrainAgentChangeRefuseVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.refuse.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 代理商拒绝改签参数
func (r *TaobaoTrainAgentChangeRefuseVtwoRequest) SetParam(_param *AgentRefuseChangeParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoTrainAgentChangeRefuseVtwoRequest) GetParam() *AgentRefuseChangeParam {
    return r._param
}
